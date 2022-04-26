import { useCallback, useEffect, useState } from "react";
import { useMounted } from "./useMounted"

export function usePromise(promise = (arg) => new Promise(), { manual = false, onSuccess = (res) => { }, onFailed = (err) => { } }, deps = null) {
  const [loading, setLoading] = useState(!manual);
  const [fetched, setFetched] = useState(manual);
  const [arg, setArg] = useState(null);

  const isMounted = useMounted();

  const fetchPromise = useCallback(() => {
    let cancel = false;

    const check = () => !cancel && isMounted();

    setFetched(true);
    promise(arg)
      .then((res) => check() && onSuccess(res, isMounted))
      .catch((err) => check() && onFailed(err, isMounted))
      .finally(() => check() && setLoading(false));

    return {
      cancel: () => { cancel = true }
    };
  }, [promise, arg, isMounted, onSuccess, onFailed]);

  useEffect(() => {
    if (!fetched && loading) {
      fetchPromise();
    }
  }, [fetched, loading, fetchPromise]);

  useEffect(() => {
    if (typeof deps === 'boolean') { }
    else if (!deps) return;
    setLoading(true);
    setFetched(false);
  }, [deps]);

  return {
    loading, fetch: (arg) => {
      setArg(arg);
      setLoading(true);
      setFetched(false);
    }
  }
}