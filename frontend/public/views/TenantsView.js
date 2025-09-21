import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Inquilinos");
  }

  async getHtml() {
    return /*html*/ `
            <h2 class="section-title">Inquilinos registrados</h2>
            <tenants-list></tenants-list>
        `;
  }
}
