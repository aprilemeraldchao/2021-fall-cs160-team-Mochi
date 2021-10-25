import SignUpWindow from "../components/signUpWindow.jsx";
import ModalWindow from "../components/modalWindow";
import Template from "../components/template.jsx";

function SignUpPage({setAuthToken}) {
  return (
    <>
      <Template body={<ModalWindow body={<SignUpWindow setAuthToken={setAuthToken}/>} />} />
    </>
  );
}

export default SignUpPage;
