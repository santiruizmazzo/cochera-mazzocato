import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="payments-view">
        <header class="section-header">
          <h2>Pagos recibidos</h2>
          <button class="new-action-btn open-modal-btn">
            <svg><use href="#person_edit"/></svg>
            Nuevo pago
          </button>
        </header>

        <div>Acá irían los pagos realizados</div>
      </section>
    `;
  }
}
