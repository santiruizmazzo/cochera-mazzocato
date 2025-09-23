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

  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  async connectedCallback() {
    if (this.version) return;

    await this.fetchVersion();
    this.render();
  }

  async fetchVersion() {
    const HEALTH_URL = import.meta.env.VITE_API_URL + "/api/health";

    await fetch(HEALTH_URL)
      .then((response) => response.json())
      .then((json) => {
        this.version = json["version"];
      })
      .catch((error) => {
        console.error("Error fetching version:", error);
      });
  }

  render() {
    const textToShow = this.version ? `v${this.version}` : "🚫";
    this.shadowRoot.querySelector("div").innerHTML = textToShow;
  }
}

customElements.define("version-box", VersionBox);
