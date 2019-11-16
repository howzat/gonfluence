package files

type Project struct {
	Name  string
	Files []MarkdownFile
}

type Projects struct {
	Projects map[string]Project
}

func (p *Projects) AddMarkdown(md *MarkdownFile) *Projects {

	if p.Projects == nil {
		p.Projects = make(map[string]Project)
	}

	project, seen := p.Projects[md.ProjectName]

	if seen {
		project.Files = append(project.Files, *md)
		p.Projects[md.ProjectName] = project
	} else {
		p.Projects[md.ProjectName] = Project{md.ProjectName, []MarkdownFile{*md}}
	}

	return p
}

func (p *Projects) Names() []string {

	var names []string
	for n := range p.Projects {
		names = append(names, n)
	}
	return names
}
