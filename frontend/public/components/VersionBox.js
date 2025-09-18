const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    .version-box {
      background-color: var(--gray-color);
      padding: 6px;
      position: relative;
      height: 30px;
      width: 60px;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  </style>

  <div class='version-box'></div>
`;

export class VersionBox extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
    this.version = null;
  }

  connectedCallback() {
    if (this.version) return;

    fetch(window.ENV.API_URL + "/api/health")
      .then((response) => response.json())
      .then((json) => {
        this.version = json["version"];
        this.shadowRoot.querySelector(".version-box").innerHTML =
          `v${this.version}`;
      })
      .catch((error) => {
        console.error("Error fetching version:", error);
        this.shadowRoot.querySelector(".version-box").innerHTML = "Error";
      });
  }
}

customElements.define("version-box", VersionBox);
