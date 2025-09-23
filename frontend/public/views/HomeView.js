import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Inicio");
  }

  async getHtml() {
    return /*html*/ `
      <h2>ESTE ES EL INICIO</h2>
    `;
  }
}
