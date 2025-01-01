package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestSetDefaults_Basic performs if function works as expected.
func TestSetDefaults_Basic(t *testing.T) {
	type ExampleConfig struct {
		FieldA string `mapstructure:"field_a" default:"value_a"`
		FieldB int    `mapstructure:"field_b" default:"42"`
	}

	v := viper.New()
	cfg := ExampleConfig{}

	SetDefaults(v, "", &cfg)

	assert.Equal(t, "value_a", v.GetString("field_a"))
	assert.Equal(t, 42, v.GetInt("field_b"))
}

// TestSetDefaults_WithPrefix test if prefix-able function also works.
func TestSetDefaults_WithPrefix(t *testing.T) {
	type NestedConfig struct {
		InnerField string `mapstructure:"inner_field" default:"nested_value"`
	}
	type ExampleConfig struct {
		Nested NestedConfig `mapstructure:"nested"`
	}

	v := viper.New()
	cfg := ExampleConfig{}

	SetDefaults(v, "prefix", &cfg)

	assert.Equal(t, "nested_value", v.GetString("prefix.nested.inner_field"))
}

// TestSetDefault_NoDefaultTags tests if empty tags are not populated.
func TestSetDefaults_NoDefaultTags(t *testing.T) {
	type ExampleConfig struct {
		FieldA string `mapstructure:"field_a"`
	}

	v := viper.New()
	cfg := ExampleConfig{}

	SetDefaults(v, "", &cfg)

	assert.Empty(t, v.GetString("field_a"))
}

// TestSetDefaults_InvalidConfigType ensures the program will not panic if no valid configuration is provided.
func TestSetDefaults_InvalidConfigType(t *testing.T) {
	v := viper.New()

	SetDefaults(v, "", "invalid")
	assert.Empty(t, v.AllKeys())
}

// TestSetDefaults_NestedStructs tests if correct handling of nesting is made.
func TestSetDefaults_NestedStructs(t *testing.T) {
	type NestedConfig struct {
		InnerField string `mapstructure:"inner_field" default:"nested_value"`
	}
	type ExampleConfig struct {
		FieldA string       `mapstructure:"field_a" default:"value_a"`
		Nested NestedConfig `mapstructure:"nested"`
	}

	v := viper.New()
	cfg := ExampleConfig{}

	SetDefaults(v, "", &cfg)

	assert.Equal(t, "value_a", v.GetString("field_a"))
	assert.Equal(t, "nested_value", v.GetString("nested.inner_field"))
}

// TestSetDefaults_Perform tests the performance of SetDefaults function.
func TestSetDefaults_Perform(t *testing.T) {

	type NestedConfig struct {
		InnerField string `mapstructure:"inner_field" default:"nested_value"`
	}
	type ExampleConfig struct {
		FieldA string       `mapstructure:"field_a" default:"value_a"`
		Nested NestedConfig `mapstructure:"nested"`
	}

	v := viper.New()
	cfg := ExampleConfig{}

	start := time.Now()
	SetDefaults(v, "", &cfg)
	duration := time.Since(start)

	durationMicroseconds := duration.Microseconds()

	if durationMicroseconds > 100000 {
		t.Errorf("Performance test failed: SetDefaults took %d microseconds, expected less than 100,000 microseconds", durationMicroseconds)
	}

}
