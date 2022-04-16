package config

import (
	"os"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_emptyEnv(t *testing.T) {
	NewEnvTest(nil).Run(func() {
		conf, err := New(nil)
		require.EqualError(t, err, "invalid config settings: No token configured")
		assert.Nil(t, conf)
	})
}

func TestNew_minimalEnv(t *testing.T) {
	NewEnvTest(map[string]string{
		"UNMS_EXPORTER_TOKEN": "a=b",
	}).Run(func() {
		conf, err := New(nil)
		require.NoError(t, err)
		assert.EqualValues(t, &Config{
			ServerAddr:   DefaultServerAddress,
			LogLevel:     DefaultLogLevel,
			TokenPerHost: tokenMap{"a": "b"},
			ExtraMetrics: nil,
		}, conf)
	})
}

func TestNew_multipleToken(t *testing.T) {
	NewEnvTest(map[string]string{
		"UNMS_EXPORTER_TOKEN": "a=b,c=d,e==f",
	}).Run(func() {
		conf, err := New(nil)
		require.NoError(t, err)
		assert.EqualValues(t, &Config{
			ServerAddr: DefaultServerAddress,
			LogLevel:   DefaultLogLevel,
			TokenPerHost: tokenMap{
				"a": "b",
				"c": "d",
				"e": "=f",
			},
			ExtraMetrics: nil,
		}, conf)
	})
}

func TestNew_flagsTakePriorityOverEnv(t *testing.T) {
	NewEnvTest(map[string]string{
		"UNMS_EXPORTER_TOKEN":  "a=b",
		"UNMS_EXPORTER_LISTEN": "[::1]:1234",
	}).Run(func() {
		conf, err := New([]string{"--listen", "[fe80::1]:9806"})
		require.NoError(t, err)
		assert.EqualValues(t, &Config{
			ServerAddr:   "[fe80::1]:9806",
			LogLevel:     DefaultLogLevel,
			TokenPerHost: tokenMap{"a": "b"},
			ExtraMetrics: nil,
		}, conf)
	})
}

func TestNew_withConfigFile(t *testing.T) {
	NewEnvTest(nil).Run(func() {
		conf, err := New([]string{"--config", "testdata/valid_config.yml"})
		require.NoError(t, err)
		assert.EqualValues(t, &Config{
			ServerAddr:   "[::1]:1234",
			LogLevel:     logrus.WarnLevel,
			TokenPerHost: tokenMap{"a.example.com": "abc"},
			ExtraMetrics: nil,
		}, conf)
	})
}

func TestNew_extraMetricsFromEnv(t *testing.T) {
	NewEnvTest(map[string]string{
		"UNMS_EXPORTER_TOKEN":         "a=b",
		"UNMS_EXPORTER_EXTRA_METRICS": "ping,ccq",
	}).Run(func() {
		conf, err := New(nil)
		require.NoError(t, err)
		assert.EqualValues(t, &Config{
			ServerAddr:   DefaultServerAddress,
			LogLevel:     DefaultLogLevel,
			TokenPerHost: tokenMap{"a": "b"},
			ExtraMetrics: []string{"ping", "ccq"},
		}, conf)
	})
}

func TestNew_extraMetricsFromSingleFlag(t *testing.T) {
	NewEnvTest(map[string]string{
		"UNMS_EXPORTER_TOKEN": "a=b",
	}).Run(func() {
		conf, err := New([]string{"--extra-metrics", "ping,ccq"})
		require.NoError(t, err)
		assert.EqualValues(t, &Config{
			ServerAddr:   DefaultServerAddress,
			LogLevel:     DefaultLogLevel,
			TokenPerHost: tokenMap{"a": "b"},
			ExtraMetrics: []string{"ping", "ccq"},
		}, conf)
	})
}

func TestNew_extraMetricsFromMultipleFlags(t *testing.T) {
	NewEnvTest(map[string]string{
		"UNMS_EXPORTER_TOKEN": "a=b",
	}).Run(func() {
		conf, err := New([]string{"--extra-metrics", "link", "--extra-metrics", "wifi"})
		require.NoError(t, err)
		assert.EqualValues(t, &Config{
			ServerAddr:   DefaultServerAddress,
			LogLevel:     DefaultLogLevel,
			TokenPerHost: tokenMap{"a": "b"},
			ExtraMetrics: []string{"link", "wifi"},
		}, conf)
	})
}

type envTest struct {
	orig map[string]string
	test map[string]string
}

func NewEnvTest(env map[string]string) *envTest {
	test := &envTest{
		orig: make(map[string]string),
		test: env,
	}
	for k := range env {
		if orig := os.Getenv(k); orig != "" {
			test.orig[k] = orig
		}
	}
	return test
}

// there can only be at most one test modifying the environment
var envTestMu sync.Mutex

func (e *envTest) Run(runner func()) {
	envTestMu.Lock()
	defer envTestMu.Unlock()

	e.Apply(e.test)       // set test env
	defer e.Clear()       // remove any test env vars
	defer e.Apply(e.orig) // reset original env

	runner()
}

func (e *envTest) Apply(env map[string]string) {
	for k, v := range env {
		os.Setenv(k, v)
	}
}

func (e *envTest) Clear() {
	for k := range e.test {
		os.Unsetenv(k)
	}
}
