#!/bin/sh

set -e

echo "$(date +"%Y-%m-%d - %H:%M:%S") | starting the entrypoint.sh"

export ORASUL_BASE_DIR="/opt/orasul" ${ORASUL_BASE_DIR}

# Execute Startup Scripts
if [ -d "${ORASUL_BASE_DIR}/entrypoint.d" ]; then

  echo "$(date +"%Y-%m-%d - %H:%M:%S") | executing scripts from [${ORASUL_BASE_DIR}/entrypoint.d]"

  find ${ORASUL_BASE_DIR}/entrypoint.d -maxdepth 1 -iname "*.sh" -type f \
    -exec /bin/sh -c "echo '########################### - {}'" \; \
    -exec /bin/sh -c "{}" \;
fi

echo "###########################"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | finished entrypoint"

exec $@ 