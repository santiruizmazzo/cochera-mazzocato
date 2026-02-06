import AbstractView from "./AbstractView.js";
import ContentSection from "../components/ContentSection.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  render() {
    const homeSection = new ContentSection();
    homeSection.title = "Pagos recibidos";
    return homeSection;
  }

  async getHtml() {
    return /*html*/ `
      <section class="payments-view">
        <custom-modal title="Registrar pago"></custom-modal>
      
        <header class="section-header">
          <h2>Pagos recibidos</h2>
          <open-modal-button>
            <svg><use href="#person_add"></svg>
            Nuevo pago
          </open-modal-button>
        </header>

        <div>Acá irían los pagos realizados</div>
        <div></div>
      </section>
    `;
  }

  setUpJavascript() {}
}
