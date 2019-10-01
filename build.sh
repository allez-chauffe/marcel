current=$(pwd)

case "${1}" in
  "all")
    echo "Building back..." && go build
    cd $current/frontend && yarn && yarn build
    cd $current/backoffice && yarn && yarn build
    ;;
  "fronts")
    cd $current/frontend && yarn && yarn build
    cd $current/backoffice && yarn && yarn build
    ;;
  "back")
    echo "Building back..." && go build
    ;;
  "")
    echo "missing argument. Usage is 'build.sh command' with command one of:\n - all\n - fronts\n - back"
esac

cd $current