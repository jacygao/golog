package golog

import (
	"reflect"
	"testing"
)

func TestGolog(t *testing.T) {
	logger := New(FATAL)

	logger.Debug("Testing logging message on DEBUG level")
	logger.Debugf("Testing logging message on DEBUG level with value %s", "mock")
	logger.Debugw("Testing logging message on DEBUG level with tags", "key", "value", "key2", "value2")
	if len(logger.tag) > 0 {
		t.Fatalf("Expected empty slice but received %v!", logger.tag)
	}

	logger.Info("Testing logging message on INFO level")
	logger.Infof("Testing logging message on INFO level with value %s", "mock")
	logger.Infow("Testing logging message on INFO level with tags", "key", "value", "key2", "value2")
	if len(logger.tag) > 0 {
		t.Fatalf("Expected empty slice but received %v!", logger.tag)
	}

	logger.Warn("Testing logging message on WARN level")
	logger.Warnf("Testing logging message on WARN level with value %s", "mock")
	logger.Warnw("Testing logging message on WARN level with tags", "key", "value", "key2", "value2")
	if len(logger.tag) > 0 {
		t.Fatalf("Expected empty slice but received %v!", logger.tag)
	}

	logger.Error("Testing logging message on ERROR level")
	logger.Errorf("Testing logging message on ERROR level with value %s", "mock")
	logger.Errorw("Testing logging message on ERROR level with tags", "key", "value", "key2", "value2")
	if len(logger.tag) > 0 {
		t.Fatalf("Expected empty slice but received %v!", logger.tag)
	}

	childLogger := logger.With("process", "testing", "program", "golog")
	childLogger.Debug("Testing child logging message on DEBUG level")
	childLogger.Info("Testing child logging message on INFO level")
	childLogger.Warn("Testing child logging message on WARN level")
	childLogger.Error("Testing child logging message on ERROR level")
	expected2 := []Tag{{"process", "testing"}, {"program", "golog"}}
	if !reflect.DeepEqual(childLogger.tag, expected2) {
		t.Fatalf("Expected %v but received %v!", expected2, childLogger.tag)
	}

	childLogger.Debugw("Testing logging message on DEBUG level with tags", "key", "value", "key2", "value2")
	if !reflect.DeepEqual(childLogger.tag, expected2) {
		t.Fatalf("Expected %v but received %v!", expected2, childLogger.tag)
	}
	childLogger.Infow("Testing logging message on INFO level with tags", "key", "value", "key2", "value2")
	if !reflect.DeepEqual(childLogger.tag, expected2) {
		t.Fatalf("Expected %v but received %v!", expected2, childLogger.tag)
	}
	childLogger.Warnw("Testing logging message on WARN level with tags", "key", "value", "key2", "value2")
	if !reflect.DeepEqual(childLogger.tag, expected2) {
		t.Fatalf("Expected %v but received %v!", expected2, childLogger.tag)
	}
	childLogger.Errorw("Testing logging message on ERROR level with tags", "key", "value", "key2", "value2")
	if !reflect.DeepEqual(childLogger.tag, expected2) {
		t.Fatalf("Expected %v but received %v!", expected2, childLogger.tag)
	}
}
