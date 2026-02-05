import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="tenant-detail-view">
        <dialog id="update-tenant-modal" closedby="any">
          <header>
            <h3>Editar inquilino</h3>
            <close-modal-button></close-modal-button>
          </header>
          <tenant-form mode="edit"></tenant-form>
        </dialog>
      
        <header class="section-header">
          <div class="section-title">
            <a href="/inquilinos" data-link>
              <svg><use href="#left_arrow" /></svg>
            </a>
            <h2>Info del inquilino</h2>
          </div>
          <open-modal-button>
            <svg><use href="#person_edit"/></svg>
            Editar
          </open-modal-button>
        </header>
        
        <tenant-full-card id="${this.params.id}"></tenant-full-card>
        <div></div>
      </section>
    `;
  }

  async setUpJavascript() {
    const view = document.querySelector(".tenant-detail-view");
    const modal = document.querySelector("#update-tenant-modal");
    const form = document.querySelector("tenant-form");

    view.addEventListener("open-modal", () => modal.showModal());
    view.addEventListener("close-modal", () => modal.close());

    view.addEventListener("tenants:updated", async (event) => {
      const tenantData = event.detail;
      tenantCard.loadTenant(tenantData);
      modal.close();
    });

    await this.fetchTenantData();
    form.loadTenant(this.tenantData);
    const tenantCard = document.querySelector("tenant-full-card");
    tenantCard.loadTenant(this.tenantData);
  }

  async fetchTenantData() {
    const TENANT_URL = `${import.meta.env.VITE_API_URL}/api/tenants/${this.params.id}`;

    await fetch(TENANT_URL)
      .then((response) => response.json())
      .then((json) => {
        this.tenantData = json;
      })
      .catch((error) => {
        console.error("Error fetching tenants:", error);
      });
  }
}
