#!/usr/bin/env bash
set -euo pipefail

# Usa: ./simulacion_crear_contenedores.sh 202412345
# Si no pasas el carnet, te lo pedirá.
CARNET="${1:-}"
if [[ -z "$CARNET" ]]; then
  read -rp "Ingresa tu carnet (sin espacios): " CARNET
  [[ -n "$CARNET" ]] || { echo "Carnet requerido"; exit 1; }
fi

# Directorio del propio script (para crear los archivos allí)
SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)"

# Cantidad aleatoria entre 1 y 4
CANT=$(( RANDOM % 4 + 1 ))

echo "Se crearán $CANT archivo(s) en: $SCRIPT_DIR"
for ((i=1; i<=CANT; i++)); do
  # Nombre aleatorio: 6 letras minúsculas
  NOMBREALEATORIO="$(head -c 100 /dev/urandom | tr -dc 'a-z' | head -c 6)"
  [[ -n "$NOMBREALEATORIO" ]] || NOMBREALEATORIO="${RANDOM}"

  NOMBRE="contenedor_${CARNET}_${NOMBREALEATORIO}.txt"
  RUTA="${SCRIPT_DIR}/${NOMBRE}"

  # Contenido = nombre del archivo
  echo "$NOMBRE" > "$RUTA"

  echo "Creado: $RUTA"
done

echo "Listo."

