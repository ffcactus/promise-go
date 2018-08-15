const resizer = document.getElementById('category-resizer');
const target = document.getElementById('category');
const parent = resizer.parentElement;

// this listener is required.
resizer.addEventListener('dragstart', ()=> {
});

parent.addEventListener('dragover', e=> {
  const parentWidth = e.currentTarget.parentElement.clientWidth;
  const targetWitdh = e.clientX * 100 / parentWidth;
  const css = targetWitdh + '%';
  target.style.flexBasis = css;
});

resizer.addEventListener('dragend', e=> {
  const parentWidth = e.currentTarget.parentElement.clientWidth;
  const targetWitdh = e.clientX * 100 / parentWidth;
  const css = targetWitdh + '%';
  target.style.flexBasis = css;
});

