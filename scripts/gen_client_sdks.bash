#!/bin/bash
# A script that generates the Livepeer AI client SDKs for several languages.
# The script uses the OpenAPI generator to generate the SDKs.
PACKAGE_NAME="livepeer_ai"
SPEC_PATH="openapi_gateway.json"
SDKS=("go" "javascript" "python" "ruby" "typescript")
OUTPUT_DIR="sdks"

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "Node.js is required to run this script. Please install Node.js and try again."
    exit 1
fi

# Generate SDKs
echo "Generating SDKs..."
for SDK in "${SDKS[@]}"; do
    echo "Generating $SDK SDK..."
    npx @openapitools/openapi-generator-cli generate -i $SPEC_PATH -g $SDK -o $OUTPUT_DIR/$SDK -p packageName=$PACKAGE_NAME
    if [ $? -eq 0 ]; then
        echo "$SDK SDK generated successfully."
    else
        echo "Failed to generate $SDK SDK."
        exit 1
    fi
done

echo "All SDKs generated successfully."
