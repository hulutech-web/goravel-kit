function T(t,i){const a=[{column:"created_at",type_name:"datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)",index:!0},{column:"updated_at",type_name:"datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)",index:!0},{column:"deleted_at",type_name:"datetime(3) DEFAULT NULL"}],m={column:"id",type_name:"bigint(20) unsigned NOT NULL AUTO_INCREMENT",primary_key:!0},n=i.map(e=>({...e,def:`${e.column} ${e.type_name}${e.not_null==="1"?" NOT NULL":""}${e.default!==void 0&&e.default!==null?` DEFAULT ${e.default}`:""}${e.unique==="1"?" UNIQUE":""}`}));n.some(e=>e.primary_key==="1")&&(console.warn('Primary key should only be set for the "id" field.'),n.forEach(e=>delete e.primary_key));const u=[`${m.column} ${m.type_name}`,...n.map(e=>e.def),...a.map(e=>`${e.column} ${e.type_name}`)].join(`,
  `),r=a.filter(e=>e.index).map(e=>`KEY idx_${t}_${e.column} (${e.column})`).join(`,
  `);return`
CREATE TABLE ${t} (
  ${u}
  , PRIMARY KEY (id)
  ${r?`, ${r}`:""}
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
`.trim().replace(/\n\s*,\s*\n/g,`
  `).replace(/,\s*$/,"")}export{T as generateCreateTableSQL};
//# sourceMappingURL=gen.4f0582cd.1755837291293.js.map
