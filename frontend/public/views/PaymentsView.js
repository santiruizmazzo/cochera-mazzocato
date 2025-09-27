import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
  }

  async getHtml() {
    return /*html*/ `
      <section class="payments-view">
        <h2>Acá se encontrarán todos los pagos de cuotas recibidos... (próximamente)</h2>
      </section>
    `;
  }
}
