package extraction

const PromptExtract = `
Analyze the following text and extract out the relevant fields into json format

Text: 
###
{{textocr}}
###

Format:
###json
{{targets}}
###

Rules:
1. If no relevant field found, omit field.
2. Use the order of the text to help you determine which text belongs in which fields. If you are unsure, omit field.
3. Array values can have multiple elements.
4. Do not include escape \n in output

JSON Output:
`
