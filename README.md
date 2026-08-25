# Cochera Mazzocato

## Versionado

La app usa un único número de versión (`X.Y.Z`, [semver](https://semver.org/)) para todo el producto,
registrado como tags de git:

- **Major**: cambios incompatibles en la API o migraciones de base de datos irreversibles.
- **Minor**: funcionalidad nueva compatible con lo existente.
- **Patch**: fixes y cambios que no agregan funcionalidad.

La versión vive en dos lugares que deben coincidir con el tag: `backend/application/version.go`
(expuesta en `GET /health`) y `frontend/package.json`.

### Hacer un release

```
make release VERSION=X.Y.Z
make push-release
```

`make release` valida que el working tree esté limpio y en `main`, actualiza ambos archivos de
versión, y crea el commit y el tag `vX.Y.Z` localmente (no pushea nada). `make push-release` publica
`main` junto con sus tags (`git push origin main --follow-tags`).

**Mergear a `main` no deploya a producción.** El deploy solo se dispara al pushear un tag `v*`, y
sólo después de que un job de verificación confirme que el tag coincide con la versión en el código
y que el build y los tests unitarios del backend pasan (ver `.github/workflows/release.yml`).