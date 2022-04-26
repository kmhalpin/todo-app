import ReactQuill from "react-quill";

const TextEditor = ({ onChange, readonly, value }) => {
  return <ReactQuill
    theme={readonly ? 'bubble' : 'snow'}
    readOnly={readonly}
    value={value}
    onChange={onChange}
  />
}

export default TextEditor;