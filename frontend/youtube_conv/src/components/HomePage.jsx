import { useState } from "react";

function HomePage() {
  const [userYTLink, setUserYTLink] = useState("");

  function getYoutubeLink(event) {
    event.preventDefault();

    const input = event.currentTarget.elements.userInputLink;
    setUserYTLink(input.value);
    console.log(input.value);
  }

  return (
    <div className="container mt-5">
      <div className="row justify-content-center">
        <div className="col-md-8">
          <form onSubmit={getYoutubeLink} className="card p-4 shadow-sm">
            <h2 className="mb-4 text-center">YouTube MP3/MP4 Converter</h2>

            <div className="mb-3">
              <label htmlFor="youtube-link" className="form-label">
                YouTube Link
              </label>
              <input
                id="youtube-link"
                type="text"
                name="userInputLink"
                className="form-control"
                placeholder="https://www.youtube.com/watch?v=..."
              />
            </div>

            <button type="submit" className="btn btn-primary w-100">
              Convert
            </button>

            {userYTLink && (
              <p className="mt-3 mb-0 text-center">
                Submitted link: {userYTLink}
              </p>
            )}
          </form>
        </div>
      </div>
    </div>
  );
}

export default HomePage;