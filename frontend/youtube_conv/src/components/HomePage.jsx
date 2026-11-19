import { useState } from "react";

function HomePage(){

    const[userYTLink,SetUserYTLink] = useState("");
    
    function getYoutubeLink(event){
        event.preventDefault();

        const input = event.currentTarget.elements.userInputLink;
    SetUserYTLink(input.value);
    console.log(userYTLink);
    

    }


    return(
        <>
            <form onSubmit={getYoutubeLink}>
                <label>Youtube Link</label>
                <input 
                    type="text" 
                    name="userInputLink"
                    defaultValue=""
                />
                <button type="submit">Link</button>
            </form>
        
        </>
    )
}


export default HomePage;