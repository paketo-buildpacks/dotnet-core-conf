package fakes

import "sync"

type ProjectParser struct {
	ASPNetIsRequiredCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Path string
		}
		Returns struct {
			Bool  bool
			Error error
		}
		Stub func(string) (bool, error)
	}
	NodeIsRequiredCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Path string
		}
		Returns struct {
			Bool  bool
			Error error
		}
		Stub func(string) (bool, error)
	}
	ParseVersionCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Glob string
		}
		Returns struct {
			String string
			Error  error
		}
		Stub func(string) (string, error)
	}
}

func (f *ProjectParser) ASPNetIsRequired(param1 string) (bool, error) {
	f.ASPNetIsRequiredCall.Lock()
	defer f.ASPNetIsRequiredCall.Unlock()
	f.ASPNetIsRequiredCall.CallCount++
	f.ASPNetIsRequiredCall.Receives.Path = param1
	if f.ASPNetIsRequiredCall.Stub != nil {
		return f.ASPNetIsRequiredCall.Stub(param1)
	}
	return f.ASPNetIsRequiredCall.Returns.Bool, f.ASPNetIsRequiredCall.Returns.Error
}
func (f *ProjectParser) NodeIsRequired(param1 string) (bool, error) {
	f.NodeIsRequiredCall.Lock()
	defer f.NodeIsRequiredCall.Unlock()
	f.NodeIsRequiredCall.CallCount++
	f.NodeIsRequiredCall.Receives.Path = param1
	if f.NodeIsRequiredCall.Stub != nil {
		return f.NodeIsRequiredCall.Stub(param1)
	}
	return f.NodeIsRequiredCall.Returns.Bool, f.NodeIsRequiredCall.Returns.Error
}
func (f *ProjectParser) ParseVersion(param1 string) (string, error) {
	f.ParseVersionCall.Lock()
	defer f.ParseVersionCall.Unlock()
	f.ParseVersionCall.CallCount++
	f.ParseVersionCall.Receives.Glob = param1
	if f.ParseVersionCall.Stub != nil {
		return f.ParseVersionCall.Stub(param1)
	}
	return f.ParseVersionCall.Returns.String, f.ParseVersionCall.Returns.Error
}
