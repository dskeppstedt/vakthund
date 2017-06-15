if sh "$(dirname -- "$0")"/deploy_test.curl | grep -q "200 OK"; then echo "\n TEST OK!";else echo "\n TEST FAIL" && exit 1; fi
if sh "$(dirname -- "$0")"/event_not_status.curl | grep -q "202 Accepted"; then echo "\n TEST OK!";else echo "\n TEST FAIL" && exit 1; fi
if sh "$(dirname -- "$0")"/status_error.curl | grep -q "202 Accepted"; then echo "\n TEST OK!";else echo "\n TEST FAIL" && exit 1; fi
if sh "$(dirname -- "$0")"/status_pending.curl | grep -q "HTTP/1.1 202 Accepted"; then echo "\n TEST OK!";else echo "\n TEST FAIL" && exit 1; fi