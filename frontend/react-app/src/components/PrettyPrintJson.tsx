import React from "react";

// Accepts any valid JSON object
interface PrettyPrintJsonProps {
  jsonData: object;
}

const PrettyPrintJson: React.FC<PrettyPrintJsonProps> = ({ jsonData }) => {
  const formattedJson = JSON.stringify(jsonData, null, 2);

  return <pre>{formattedJson}</pre>;
};

export default PrettyPrintJson;