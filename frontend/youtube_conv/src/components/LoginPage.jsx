import { useState } from "react"


function LoginPage(){

    const[email,SetEmail]=useState("");
    const[password,SetPassword]=useState("");

    function HandleSubmit(event){
        event.preventDefault();

        console.log(email);
        console.log(password);
        
    }

    return(
        <>
            <form onSubmit={HandleSubmit}>
                <label>Email</label>
                <input 
                    type="email" 
                    name="user_email"
                    placeholder="Email"
                    value={email}
                    onChange={(event)=>SetEmail(event.target.value)} // callback function () => seteamail()
                /> 
                <br />
                <label>Password</label>
                <input 
                    type="text" 
                    name="user_passward" 
                    placeholder="password"
                    value={password}
                    onChange={(event)=>SetPassword(event.target.value)}
                /> 
                <br />

                <button type="submit">Login</button>
            </form>
        
        </>

    )
}


export default LoginPage