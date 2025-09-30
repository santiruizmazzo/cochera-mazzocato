const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    div {
      display: none;
      padding: 0.3rem;
      border: 0.125rem solid red;
      background-color: rgba(255, 195, 195, 1);
      color: red;
    }
  </style>

  <div></div>
`;

export default class ErrorBox extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  show(message) {
    const box = this.shadowRoot.querySelector("div");
    box.style.display = "block";
    const capitalizedMessage =
      message.charAt(0).toUpperCase() + message.slice(1);
    box.innerHTML = capitalizedMessage;
  }

  hide() {
    const box = this.shadowRoot.querySelector("div");
    box.style.display = "none";
    box.innerHTML = "";
  }
}

customElements.define("error-box", ErrorBox);
