import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Pagos");
  }

  async getHtml() {
    return /*html*/ `
      <h2>AQUI ESTARIAN LOS PAGOS RECIBIDOS</h2>
    `;
  }
}
