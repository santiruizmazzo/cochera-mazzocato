import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Inicio");
  }

  async getHtml() {
    return /*html*/ `
      <svg class="spinner"><use href="#spinner"/></svg>
    `;
  }
}
