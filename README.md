Tutorial project for article in [Hashnode.dev](https://skylab.hashnode.dev/operator-sdk-beginner-tutorial-ckqv88l3o03tglls19m8jefb1)

### How to run
1. Check you have access to kubenetes cluster via `kubectl`, check by running `kubectl get nodes` or `kubectl get pods`.

2. Run `make install run` to install CRDs and run controller.

3. Go to `example` folder and run `kubectl apply -f .` to apply all manifests in it.

4. Check output one by one using `kubectl get add`, `kubectl get subs`, `kubectl get mul`, `kubectl get div` or get all as one `kubectl get add,subs,mul,div`.