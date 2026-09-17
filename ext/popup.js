const repo = 'bonsai/LAB-Rapid-Prototype';
const runs = document.querySelector('#runs');
async function load() {
  runs.innerHTML = '<li>Loading…</li>';
  const r = await fetch(`https://api.github.com/repos/${repo}/actions/runs?per_page=8`);
  if (!r.ok) throw new Error(`${r.status} ${r.statusText}`);
  const data = await r.json();
  runs.innerHTML = data.workflow_runs.map(x => {
    const cls = x.conclusion === 'success' ? 'ok' : x.conclusion === 'failure' ? 'bad' : '';
    return `<li class="${cls}"><a target="_blank" href="${x.html_url}">${x.name}</a><br>${x.status} / ${x.conclusion ?? 'running'}</li>`;
  }).join('') || '<li>No runs</li>';
}
document.querySelector('#refresh').onclick = load;
load().catch(e => runs.innerHTML = `<li class="bad">${e}</li>`);
