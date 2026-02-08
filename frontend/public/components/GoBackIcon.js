const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    * {
      box-sizing: border-box;
    }

    svg {
      width: 1.7rem;
      height: 1.7rem;

      &:hover {
        cursor: pointer;
      }
    }
  </style>

  <svg viewBox="0 -960 960 960">
    <path d="M400-240 160-480l240-240 56 58-142 142h486v80H314l142 142z"/>
  </svg>
`;

export default class GoBackIcon extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
    this.addEventListener("click", () => history.back());
  }
}

customElements.define("go-back-icon", GoBackIcon);
