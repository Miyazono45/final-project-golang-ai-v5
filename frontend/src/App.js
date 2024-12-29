import React, { useState } from "react";
import axios from "axios";
import "./App.css";
import "./";

function App() {
  const [file, setFile] = useState(null);
  const [initialQuery, setInitialQuery] = useState("");
  const [query, setQuery] = useState("");
  const [responseChatbot, setResponseChatbot] = useState("");
  const [response, setResponse] = useState("");

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleUpload = async () => {
    const formData = new FormData();
    formData.append("file", file);
    formData.append("initial_query", initialQuery);

    try {
      const res = await axios.post("http://localhost:8080/upload", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });
      const parsedJSON = JSON.parse(res.data);
      console.log(parsedJSON);
      setResponseChatbot(parsedJSON.answer); // Assuming the response has an 'answer' field
    } catch (error) {
      console.error("Error uploading file:", error);
      setResponseChatbot(error.response.data);
    }
  };

  const handleChat = async () => {
    const formData = new FormData();
    formData.append("query", query);

    try {
      const res = await axios.post("http://localhost:8080/chat", formData);
      // const parsedJSON = JSON.parse(res.data.answer);
      console.log(res.data.answer);
      setResponse(res.data.answer);
    } catch (error) {
      console.error("Error querying chat:", error);
    }
  };

  return (
    <div className="overflow-hidden font-rubik w-[95vw] border-x-[1px] border-[#151715] translate-x-[2.5vw] relative">
      <div className="w-screen h-screen relative overflow-hidden">
        {/* Kiri */}
        {/* Text Only */}
        <div className="absolute z-30 left-[4vw] top-[25vh]">
          <div className="w-full h-full flex justify-center flex-col mb-[8vh]">
            <div className="">
              <h1 className="text-[4rem] font-bold bg-clip-text text-transparent bg-gradient-to-r from-[#64DA5A] to-[#7BFA6F] tracking-tig drop-shadow-lg">
                Data Analysis Chatbot
              </h1>

              <h1 className="text-[4rem] font-bold tracking-tight drop-shadow-lg -mt-6 w-fit text-transparent bg-clip-text bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500">
                & Mini GPT!
              </h1>
            </div>

            <div>
              <p className="text-[#0B3807] font-medium text-xl">
                Upload, Ask, Know!
              </p>
              <p className="text-[#0B3807] font-light text-lg">
                An platform designed to provide instant data analysis solutions
                through AI-based chatbots. Get deep and accurate <br /> insights
                from your data just by typing in simple questions.
              </p>
            </div>
          </div>
        </div>
        {/* Button Only */}
        <div className="absolute w-1/2 h-1/2 left-[4vw] top-[60vh] flex flex-row gap-x-5 z-40">
          <a
            type="button"
            href="#chatbot"
            className="w-[12vw] h-[5vh] bg-[#27AA1B] text-[#F1FCF0] rounded-3xl hover:scale-110 transition-all duration-150 flex justify-center items-center"
          >
            Ask Now!
          </a>
          <a
            href="#minigpt"
            type="button"
            // value="just my linkedin"
            className="w-[12vw] h-[5vh] bg-[#fcfbff] text-purple-400 rounded-3xl border-[1px] border-[#5a3ef7] hover:scale-110 transition-all duration-150 flex justify-center items-center"
          >
            Lemme use MiniGPT
          </a>
        </div>
        {/* KANAN */}
        <div className="w-1/2 h-[80vh] flex justify-center items-center absolute right-[1.5vw] top-1/2 -translate-y-1/2 z-20">
          <img
            src="images/frontHero_2.png"
            alt="err"
            width={780}
            height={780}
            className="max-w-[780px] max-h-[780px] absolute z-10 brightness-[1.25] contrast-[.95] drop-shadow-2xl top-[-5vh]"
          ></img>
        </div>
        {/* bg nig */}
        <div className="w-screen h-screen absolute">
          <img
            src="images/irisdecent_1.jpg"
            alt=""
            className="brightness-[1.15] opacity-30 blur-lg hue-rotate-[265deg]"
            width={1920}
            height={1080}
          />
        </div>
      </div>
      {/* bg nig */}
      {/* <div className="w-screen h-screen absolute">
        <img
          src="images/irisdecent_1.jpg"
          alt=""
          className="brightness-[1.15] opacity-15 blur-lg"
          width={1920}
          height={1080}
        />
      </div> */}
      {/* Main */}
      <div className="relative mb-[15vh]" id="chatbot">
        {/* Data Analyst chatbot */}
        <div className="flex">
          {/* left */}
          <div className="w-1/2 relative left-[4vw]">
            <div className="mb-[10vh] w-full">
              <h1 className="text-[3rem] font-bold bg-clip-text text-transparent bg-gradient-to-r from-[#64DA5A] to-[#7BFA6F] tracking-tig drop-shadow-lg">
                Data Analysis <span className="text-[#151715]">Chatbot</span>
              </h1>
              <p className="text-[#0B3807] font-normal text-base mt-2">
                You can upload your CSV and give an question about your CSV!
                It's like asking "What about my Electricity?" <br /> or "What
                the Least Energy from my Data". Just use your creativity LMAO.
              </p>
              <div className="flex flex-col">
                <input
                  type="file"
                  onChange={handleFileChange}
                  className="mt-8 rounded-md w-[65%] border-[1px] border-[#174613] p-2 text-[#174613]"
                />
                <div className="flex flex-row w-full items-center mt-3">
                  <input
                    type="text"
                    value={initialQuery}
                    onChange={(e) => setInitialQuery(e.target.value)}
                    placeholder="Ask a Initial Question..."
                    className="rounded-md w-[65%] h-[50px] border-[1px] border-[#27AA1B] p-2 text-[#25901B]"
                  />
                  <button
                    onClick={handleUpload}
                    className="px-[10px] bg-[#27AA1B] text-[#F1FCF0] ml-4 border-none rounded-md w-[25%] h-[50px]"
                  >
                    Upload and Analyze
                  </button>
                </div>
              </div>
            </div>
          </div>
          {/* right */}
          <div className="w-[40vw] relative left-16 right-5">
            <div className="mt-[20px] p-[10px] border-[1px] h-[26vh] border-solid border-[#25901B] rounded-md bg-[$f9f9f9]">
              <h2 className="text-[#0B3807] font-medium text-xl">
                Chatbot Response
              </h2>
              <p className="text-[#0B3807] font-light text-base">
                {responseChatbot}
              </p>
            </div>
          </div>
        </div>

        {/* Microsoft Mini */}
        <div className="relative z-30" id="minigpt">
          <div className="flex">
            {/* Left */}
            <div className="w-1/2 relative left-[4vw]">
              <div>
                <h1 className="text-[3rem] font-bold bg-clip-text text-transparent bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 tracking-tig drop-shadow-lg">
                  Microsoft <span className="text-[#151715]">Mini 3.5</span>
                </h1>
                <p className="text-[#0B3807] font-normal text-base mt-2 w-[40vw]">
                  Hey Look! We have smoolll model GPT! (Actually it belongs to
                  microsoft){" "}
                  <a
                    href="https://huggingface.co/microsoft/Phi-3.5-mini-instruct"
                    target="_blank"
                    className="text-purple-500"
                  >
                    You can check this model in here!{" "}
                  </a>
                  Again, Use your creativity to ask something to this model.
                  Please don't bring feelings when giving commands to this
                  model.
                </p>
                <div className="flex flex-col">
                  <div className="flex flex-row w-full items-center mt-3">
                    <input
                      type="text"
                      value={query}
                      onChange={(e) => setQuery(e.target.value)}
                      placeholder="Ask a question..."
                      className="rounded-md w-[65%] h-[50px] border-[1px] border-pink-500 p-2 text-pink-500"
                    />
                    <button
                      onClick={handleChat}
                      className="px-[10px] bg-gradient-to-r from-indigo-400 via-purple-400 to-pink-400 text-white ml-4 border-none rounded-md w-[25%] h-[50px]"
                    >
                      Chat Now!
                    </button>
                  </div>
                </div>
              </div>
            </div>
            {/* right */}
            <div className="w-[40vw] relative left-16 right-5">
              <div className="mt-[20px] p-[10px] border-[1px] h-[26vh] border-solid border-indigo-500 rounded-md bg-[$f9f9f9] overflow-y-scroll">
                <h2 className="text-[#211257] font-medium text-xl">
                  Microsoft Mini Model 3.5 Response
                </h2>
                <p className="text-[#211257] font-light text-base">
                  {response}
                </p>
              </div>
            </div>
          </div>
        </div>
        {/* bg nig */}
        <div className="w-screen h-screen absolute top-[0vh] -z-10">
          <img
            src="images/irisdecent_1.jpg"
            alt=""
            className="brightness-[1.15] opacity-20 blur-lg hue-rotate-[265deg]"
            width={1920}
            height={1080}
          />
        </div>
      </div>
    </div>
  );
}

export default App;
