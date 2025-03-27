import { fileURLToPath } from "url";
import path from "path";
import fs from "fs";

// Workaround for __dirname in ES modules
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const argv2 = process.argv[2];
const componentName = argv2.charAt(0).toUpperCase() + argv2.slice(1).toLowerCase();


if (!componentName) {
  console.error("❌ Please provide a component name.");
  process.exit(1);
}

const componentDir = path.join(__dirname, "src", "components", componentName);

// Check if the folder already exists
if (fs.existsSync(componentDir)) {
  console.error("❌ Component already exists.");
  process.exit(1);
}

// Create the component folder
fs.mkdirSync(componentDir, { recursive: true });

// Create index.jsx file
const indexContent = `export { default as ${componentName} } from "./${componentName}";\n`;
fs.writeFileSync(path.join(componentDir, "index.jsx"), indexContent);

// Create the component file
const componentContent = `import styles from "./style.module.scss";

export default function ${componentName} (){
    return(
    <div>${componentName} component</div>
    )
};
`;
fs.writeFileSync(path.join(componentDir, `${componentName}.jsx`), componentContent);

// Create the CSS module file
fs.writeFileSync(path.join(componentDir, "style.module.scss"), `@use '../../global.scss' as *;`);

console.log(`✅ Component "${componentName}" created successfully in src/components/${componentName}/`);
