package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

type TestSuite struct{}

func (s TestSuite) parse(t *testing.T, a string) *html.Node {
	n, err := html.Parse(strings.NewReader(a))
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
		return nil
	}
	return n.FirstChild.FirstChild.NextSibling.FirstChild
}

func TestExtractHref(t *testing.T) {
	testCases := []struct {
		name string
		a    string
		href string
	}{
		{
			name: "valid",
			a:    `<a href="/Login">Login</a>`,
			href: "/Login",
		},
		{
			name: "missing href",
			a:    `<a>Login</a>`,
			href: "",
		},
	}

	s := TestSuite{}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			a := s.parse(t, c.a)
			href := extractHref(a)
			if href != c.href {
				t.Fatalf("Expected %s, got %s", c.href, href)
			}
		})
	}
}

func TestExtractText(t *testing.T) {
	testCases := []struct {
		name string
		a    string
		text string
	}{
		{
			name: "valid",
			a:    `<a href="/Login">Login</a>`,
			text: "Login",
		},
		{
			name: "valid: nested elements",
			a:    `<a href="/Login">Login <span>as <strong>Godazz</strong></span></a>`,
			text: "Login as Godazz",
		},
		{
			name: "valid: nested comments",
			a:    `<a href="/Login">Login <!-- This is a comment --></a>`,
			text: "Login",
		},
	}

	s := TestSuite{}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			a := s.parse(t, c.a)
			text := extractText(a)
			if text != c.text {
				t.Fatalf("Expected %s, got %s", c.text, text)
			}
		})
	}
}
