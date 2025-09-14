const template = document.createElement("template");
template.innerHTML = `
    <style>
        .version-box {
            background-color: rgb(53, 53, 53);
            padding: 6px;
            color: #fff;
        }
    </style>

    <div class='version-box'></div>
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
        this.root.querySelector(".version-box").innerHTML = this.version;
      })
      .catch((error) => {
        console.error("Error fetching version:", error);
        this.root.querySelector(".version-box").innerHTML = "Error";
      });
  }
}

customElements.define("version-box", VersionBox);
