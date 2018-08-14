const resizer = document.getElementById('category-resizer');
const target = document.getElementById('category');

// resizer.addEventListener('mouseenter', e =>{
//   e.currentTarget.style.cssText = 'flex-basis: 10px';
// });

// resizer.addEventListener('mouseout', e => {
//   e.currentTarget.style.cssText = 'flex-basis: ;';
// });

resizer.addEventListener('dragstart', e=> {
  // const parentWidth = e.currentTarget.parentElement.clientWidth;
  // const targetWitdh = e.clientX / parentWidth;
  // target.style.cssText = 'flex-basis: ' + targetWitdh + ';';
  // console.info(targetWitdh);
});

resizer.addEventListener('dragend', e=> {
  const parentWidth = e.currentTarget.parentElement.clientWidth;
  const targetWitdh = e.clientX * 100 / parentWidth;
  const css = 'flex-basis: ' + targetWitdh + '%;';
  target.style.cssText = css;
  console.info(targetWitdh);
});

