import { useState } from "react";
import { useNavigate } from "react-router-dom";

function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const [err,SetErr] = useState("");

  // debugger
  function checkValidData(){
    if(!email.includes("@") || !email.includes(".")){
      SetErr("mail does not contain the valid format");
      return false;
    }
    if(password.length < 1 && password.length < 10){
      SetErr("Password length is small or max");
      return false;
    }
  }

  async function handleSubmit(event) {
    event.preventDefault();
    if(checkValidData()){
      console.log("Data is not right :(");
      return;
    }

    console.log(email);
    console.log(password);
    debugger
    try{
    const response = await fetch("http://localhost:8080/login",{
      method: "POST",
      headers: {
        "Content-Type":"application/json",
      },
      body: JSON.stringify({
      email: email,
      password: password,
    }),
    });

    const data = await response.json();
    console.log(data);
  }catch(err){
    console.log(err);
  }
    navigate("/home");


  }

  return (
    <div className="container mt-5">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <form onSubmit={handleSubmit} className="card p-4 shadow-sm">
            <h2 className="mb-4 text-center">Login</h2>
            <div className="mb-3">
              <label htmlFor="email" className="form-label">
                Email
              </label>
              <input
                id="email"
                type="email"
                className="form-control"
                value={email}
                onChange={(event) => setEmail(event.target.value)}
              />
              {err &&
                <p>
                    oops! enter valid mail
                </p>}
            </div>

            <div className="mb-3">
              <label htmlFor="password" className="form-label">
                Password
              </label>
              <input
                id="password"
                type="password"
                className="form-control"
                value={password}
                onChange={(event) => setPassword(event.target.value)}
              />
              {err && 
              <p>
                oops! enter small range password
              </p>  
              }
            </div>

            <button type="submit" className="btn btn-primary w-100">
              Login
            </button>
          </form>
        </div>
      </div>
    </div>
  );
}

export default LoginPage;