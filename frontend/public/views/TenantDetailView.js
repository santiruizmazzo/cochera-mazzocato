import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="tenant-detail-view">
        <header class="section-header">
          <div class="section-title">
            <a href="/inquilinos" data-link>
              <svg><use href="#left_arrow" /></svg>
            </a>
            <h2>Info del inquilino</h2>
          </div>
        </header>
        <tenant-full-card id="${this.params.id}"></tenant-full-card>
        <div></div>
      </section>
    `;
  }
}
