// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jcs/id3-go (interfaces: Tagger)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "github.com/jcs/id3-go/v2"
)

// MockTagger is a mock of Tagger interface.
type MockTagger struct {
	ctrl     *gomock.Controller
	recorder *MockTaggerMockRecorder
}

// MockTaggerMockRecorder is the mock recorder for MockTagger.
type MockTaggerMockRecorder struct {
	mock *MockTagger
}

// NewMockTagger creates a new mock instance.
func NewMockTagger(ctrl *gomock.Controller) *MockTagger {
	mock := &MockTagger{ctrl: ctrl}
	mock.recorder = &MockTaggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagger) EXPECT() *MockTaggerMockRecorder {
	return m.recorder
}

// AddFrames mocks base method.
func (m *MockTagger) AddFrames(arg0 ...v2.Framer) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddFrames", varargs...)
}

// AddFrames indicates an expected call of AddFrames.
func (mr *MockTaggerMockRecorder) AddFrames(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFrames", reflect.TypeOf((*MockTagger)(nil).AddFrames), arg0...)
}

// Album mocks base method.
func (m *MockTagger) Album() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Album")
	ret0, _ := ret[0].(string)
	return ret0
}

// Album indicates an expected call of Album.
func (mr *MockTaggerMockRecorder) Album() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Album", reflect.TypeOf((*MockTagger)(nil).Album))
}

// AllFrames mocks base method.
func (m *MockTagger) AllFrames() []v2.Framer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllFrames")
	ret0, _ := ret[0].([]v2.Framer)
	return ret0
}

// AllFrames indicates an expected call of AllFrames.
func (mr *MockTaggerMockRecorder) AllFrames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllFrames", reflect.TypeOf((*MockTagger)(nil).AllFrames))
}

// Artist mocks base method.
func (m *MockTagger) Artist() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Artist")
	ret0, _ := ret[0].(string)
	return ret0
}

// Artist indicates an expected call of Artist.
func (mr *MockTaggerMockRecorder) Artist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Artist", reflect.TypeOf((*MockTagger)(nil).Artist))
}

// Bytes mocks base method.
func (m *MockTagger) Bytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes.
func (mr *MockTaggerMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockTagger)(nil).Bytes))
}

// Comments mocks base method.
func (m *MockTagger) Comments() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Comments")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Comments indicates an expected call of Comments.
func (mr *MockTaggerMockRecorder) Comments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Comments", reflect.TypeOf((*MockTagger)(nil).Comments))
}

// DeleteFrame mocks base method.
func (m *MockTagger) DeleteFrame(arg0 v2.Framer) []v2.Framer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFrame", arg0)
	ret0, _ := ret[0].([]v2.Framer)
	return ret0
}

// DeleteFrame indicates an expected call of DeleteFrame.
func (mr *MockTaggerMockRecorder) DeleteFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFrame", reflect.TypeOf((*MockTagger)(nil).DeleteFrame), arg0)
}

// DeleteFrames mocks base method.
func (m *MockTagger) DeleteFrames(arg0 string) []v2.Framer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFrames", arg0)
	ret0, _ := ret[0].([]v2.Framer)
	return ret0
}

// DeleteFrames indicates an expected call of DeleteFrames.
func (mr *MockTaggerMockRecorder) DeleteFrames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFrames", reflect.TypeOf((*MockTagger)(nil).DeleteFrames), arg0)
}

// Dirty mocks base method.
func (m *MockTagger) Dirty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dirty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Dirty indicates an expected call of Dirty.
func (mr *MockTaggerMockRecorder) Dirty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dirty", reflect.TypeOf((*MockTagger)(nil).Dirty))
}

// Frame mocks base method.
func (m *MockTagger) Frame(arg0 string) v2.Framer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Frame", arg0)
	ret0, _ := ret[0].(v2.Framer)
	return ret0
}

// Frame indicates an expected call of Frame.
func (mr *MockTaggerMockRecorder) Frame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Frame", reflect.TypeOf((*MockTagger)(nil).Frame), arg0)
}

// Frames mocks base method.
func (m *MockTagger) Frames(arg0 string) []v2.Framer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Frames", arg0)
	ret0, _ := ret[0].([]v2.Framer)
	return ret0
}

// Frames indicates an expected call of Frames.
func (mr *MockTaggerMockRecorder) Frames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Frames", reflect.TypeOf((*MockTagger)(nil).Frames), arg0)
}

// Genre mocks base method.
func (m *MockTagger) Genre() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genre")
	ret0, _ := ret[0].(string)
	return ret0
}

// Genre indicates an expected call of Genre.
func (mr *MockTaggerMockRecorder) Genre() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genre", reflect.TypeOf((*MockTagger)(nil).Genre))
}

// Padding mocks base method.
func (m *MockTagger) Padding() uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Padding")
	ret0, _ := ret[0].(uint)
	return ret0
}

// Padding indicates an expected call of Padding.
func (mr *MockTaggerMockRecorder) Padding() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Padding", reflect.TypeOf((*MockTagger)(nil).Padding))
}

// SetAlbum mocks base method.
func (m *MockTagger) SetAlbum(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAlbum", arg0)
}

// SetAlbum indicates an expected call of SetAlbum.
func (mr *MockTaggerMockRecorder) SetAlbum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAlbum", reflect.TypeOf((*MockTagger)(nil).SetAlbum), arg0)
}

// SetArtist mocks base method.
func (m *MockTagger) SetArtist(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetArtist", arg0)
}

// SetArtist indicates an expected call of SetArtist.
func (mr *MockTaggerMockRecorder) SetArtist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArtist", reflect.TypeOf((*MockTagger)(nil).SetArtist), arg0)
}

// SetGenre mocks base method.
func (m *MockTagger) SetGenre(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGenre", arg0)
}

// SetGenre indicates an expected call of SetGenre.
func (mr *MockTaggerMockRecorder) SetGenre(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGenre", reflect.TypeOf((*MockTagger)(nil).SetGenre), arg0)
}

// SetTitle mocks base method.
func (m *MockTagger) SetTitle(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTitle", arg0)
}

// SetTitle indicates an expected call of SetTitle.
func (mr *MockTaggerMockRecorder) SetTitle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTitle", reflect.TypeOf((*MockTagger)(nil).SetTitle), arg0)
}

// SetYear mocks base method.
func (m *MockTagger) SetYear(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetYear", arg0)
}

// SetYear indicates an expected call of SetYear.
func (mr *MockTaggerMockRecorder) SetYear(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetYear", reflect.TypeOf((*MockTagger)(nil).SetYear), arg0)
}

// Size mocks base method.
func (m *MockTagger) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockTaggerMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockTagger)(nil).Size))
}

// Title mocks base method.
func (m *MockTagger) Title() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Title")
	ret0, _ := ret[0].(string)
	return ret0
}

// Title indicates an expected call of Title.
func (mr *MockTaggerMockRecorder) Title() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Title", reflect.TypeOf((*MockTagger)(nil).Title))
}

// Version mocks base method.
func (m *MockTagger) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockTaggerMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockTagger)(nil).Version))
}

// Year mocks base method.
func (m *MockTagger) Year() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Year")
	ret0, _ := ret[0].(string)
	return ret0
}

// Year indicates an expected call of Year.
func (mr *MockTaggerMockRecorder) Year() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Year", reflect.TypeOf((*MockTagger)(nil).Year))
}
