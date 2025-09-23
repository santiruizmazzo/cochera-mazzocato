const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    div {
      --font-size: 0.85rem;
      --padding-relation: 0.65;

      background-color: var(--clr-bg-light);
      padding: calc(var(--font-size) * var(--padding-relation));
      position: relative;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: var(--font-size);
    }
  </style>

  <div></div>
`;

export default class VersionBox extends HTMLElement {
  version = null;
  HEALTH_ENDPOINT = import.meta.env.VITE_API_URL + "/api/health";

  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  connectedCallback() {
    if (this.version) return;

    fetch(this.HEALTH_ENDPOINT)
      .then((response) => response.json())
      .then((json) => {
        this.version = json["version"];
        this.shadowRoot.querySelector("div").innerHTML = `v${this.version}`;
      })
      .catch((error) => {
        console.error("Error fetching version:", error);
        this.shadowRoot.querySelector("div").innerHTML = "Error";
      });
  }
}

customElements.define("version-box", VersionBox);
