docker compose up -d
docker compose exec -d api make -C backend run
docker compose exec frontend bash