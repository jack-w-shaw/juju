run_appdata_basic() {
	echo

	file="${TEST_DIR}/appdata-basic.log"

	ensure "appdata-basic" "${file}"

	juju deploy juju-qa-appdata-source --series bionic
	juju deploy -n 2 juju-qa-appdata-sink --series bionic
	juju relate appdata-source appdata-sink

	wait_for "blocked" "$(workload_status appdata-source 0).current"

	juju config appdata-source token=test-value
	output=$(juju config appdata-source)
	expected=$(
		cat <<'EOF'
application: appdata-source
application-config:
  trust:
    default: false
    description: Does this application have access to trusted credentials
    source: default
    type: bool
    value: false
charm: appdata-source
settings:
  token:
    default: ""
    description: Token used to prove that the relation is working
    source: user
    type: string
    value: test-value
EOF
	)
	if [[ ${output} != "${expected}" ]]; then
		echo "expected ${expected}, got ${output}"
		exit 1
	fi

	# Wait for the token to arrive on each of the sink units.
	wait_for "test-value" "$(workload_status appdata-sink 0).message"
	wait_for "test-value" "$(workload_status appdata-sink 1).message"

	# Check that the token is in /var/run/appdata-sink/token on each
	# one.
	output=$(juju ssh appdata-sink/0 cat /var/run/appdata-sink/token)
	check_contains "$output" "appdata-source/0 test-value"

	output=$(juju ssh appdata-sink/1 cat /var/run/appdata-sink/token)
	check_contains "$output" "appdata-source/0 test-value"

	juju add-unit appdata-source
	juju remove-unit appdata-source/0

	wait_for "idle" "$(agent_status appdata-source 1).current"
	juju config appdata-source token=value2

	wait_for "value2" "$(workload_status appdata-sink 0).message"
	wait_for "value2" "$(workload_status appdata-sink 1).message"

	output=$(juju ssh appdata-sink/0 cat /var/run/appdata-sink/token)
	check_contains "$output" "appdata-source/1 value2"

	output=$(juju ssh appdata-sink/1 cat /var/run/appdata-sink/token)
	check_contains "$output" "appdata-source/1 value2"

	juju config appdata-source --reset token
	output=$(juju config appdata-source)
	expected=$(
		cat <<'EOF'
application: appdata-source
application-config:
  trust:
    default: false
    description: Does this application have access to trusted credentials
    source: default
    type: bool
    value: false
charm: appdata-source
settings:
  token:
    default: ""
    description: Token used to prove that the relation is working
    source: default
    type: string
    value: ""
EOF
	)
	if [[ ${output} != "${expected}" ]]; then
		echo "expected ${expected}, got ${output}"
		exit 1
	fi

	destroy_model "appdata-basic"
}

test_appdata_int() {
	if [ "$(skip 'test_appdata_int')" ]; then
		echo "==> TEST SKIPPED: appdata int tests"
		return
	fi

	(
		set_verbosity

		cd suites/appdata || exit

		run "run_appdata_basic"
	)
}
