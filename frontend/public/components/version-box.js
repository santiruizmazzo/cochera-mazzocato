const template = document.createElement("template");
template.innerHTML = `
    <style>
        .version-box {
            background-color: rgb(234, 234, 234, 1);
            padding: 6px;
            position: relative;
            height: 30px;
            width: 60px;
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .loading-placeholder {
            position: absolute;
            left: 0;
            top: 0;
            height: 100%;
            width: 75%;
            background-image: linear-gradient(to left,rgba(234, 234, 234, .05), rgba(211, 211, 211, 0.3), rgba(175, 175, 175, 0.6), rgba(211, 211, 211, 0.3), rgba(234, 234, 234, .05));
            animation: loading 1.5s ease-out infinite;
        }
        
        @keyframes loading {
            0%{
                left: -20%;
            }
            100%{
                left: 60%;
            }
        }
    </style>

    <div class='version-box'>
      <div class='loading-placeholder'></div>
    </div>
`;

export class VersionBox extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "closed" });
    this.root.append(template.content.cloneNode(true));
    this.version = null;
  }

  connectedCallback() {
    if (this.version) return;

    fetch(window.ENV.API_URL + "/api/health")
      .then((response) => response.json())
      .then((json) => {
        this.version = json["version"];
        this.root.querySelector(".version-box").innerHTML = `v${this.version}`;
      })
      .catch((error) => {
        console.error("Error fetching version:", error);
        this.root.querySelector(".version-box").innerHTML = "Error";
      });
  }
}

customElements.define("version-box", VersionBox);
