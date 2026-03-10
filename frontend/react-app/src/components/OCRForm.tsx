import React, { useState } from "react";
import axios from "axios";
import { OCRResponse } from "../types/OCRResponse";

const OCRForm: React.FC = () => {
  const [file, setFile] = useState<File | null>(null);
  const [result, setResult] = useState<OCRResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const API_URL = process.env.REACT_APP_BACKEND_URL || "http://localhost:8080";

  console.log(process.env.REACT_APP_BACKEND_URL)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      setFile(e.target.files[0]);
      setResult(null);
      setError(null);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!file) return;

    setLoading(true);
    setError(null);

    const formData = new FormData();
    formData.append("image", file);

    try {
      const response = await axios.post<OCRResponse>(
        `${API_URL}/ocr`,
        formData,
      );
      setResult(response.data);
    } catch (err) {
      console.error(err);
      setError("OCR failed. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          OCR Uploader
        </h2>
        <p className="mt-2 text-center text-sm text-gray-600">
          Upload an image and extract text.
        </p>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div className="bg-white py-8 px-6 shadow rounded-lg sm:px-10">
          <form className="mb-4 space-y-6" onSubmit={handleSubmit}>
            <div>
              <label
                htmlFor="file"
                className="block text-sm font-medium text-gray-700"
              >
                Select file
              </label>
              <div className="mt-1">
                <input
                  id="file"
                  name="file"
                  type="file"
                  accept="image/png, image/jpeg, application/pdf"
                  onChange={handleChange}
                  className="block w-full text-sm text-gray-900 border border-gray-300 rounded-md cursor-pointer focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>

            <div>
              <button
                type="submit"
                disabled={loading || !file}
                className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
              >
                {loading ? "Processing..." : "Upload & Run OCR"}
              </button>
            </div>
          </form>

          {error && (
            <p className="text-red-600 text-center font-medium">{error}</p>
          )}

          {result && (
            <div className="mt-6">
              <h3 className="text-lg font-medium text-gray-900 mb-2">
                OCR Result
              </h3>
              <pre className="bg-gray-100 p-4 rounded text-sm text-gray-800 whitespace-pre-wrap overflow-x-auto">
                {result?.text}
              </pre>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default OCRForm;
