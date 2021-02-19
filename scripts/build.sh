current=$(pwd)

case "${1}" in
  "all")
    echo "Building back..." && go build ./cmd/marcel
    cd $current/pkg/frontend && yarn && yarn build
    cd $current/pkg/backoffice && yarn && yarn build
    ;;
  "fronts")
    cd $current/pkg/frontend && yarn && yarn build
    cd $current/pkg/backoffice && yarn && yarn build
    ;;
  "back")
    echo "Building back..." && go build ./cmd/marcel
    ;;
  "")
    echo "missing argument. Usage is 'build.sh command' with command one of:\n - all\n - fronts\n - back"
esac

cd $current
