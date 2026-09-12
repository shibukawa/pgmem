package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v8 = int32(8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(16)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v15 = v7<<(uint(v8)%32) | v10<<(uint(v11)%32) | v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v25 = v17<<(uint(v8)%32) | v20<<(uint(v11)%32) | v24
	if base.Ui32(v15) < base.Ui32(v25) {
		v52 = int32(-1)
	} else {
		if base.Ui32(v25) < base.Ui32(v15) {
			v52 = int32(1)
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
			v32 = int32(8)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
			v35 = int32(16)
			v38 = v30 | (v31<<(uint(v32)%32) | v34<<(uint(v35)%32))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)))
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
			v47 = v39 | (v40<<(uint(v32)%32) | v43<<(uint(v35)%32))
			if base.Ui32(v38) < base.Ui32(v47) {
				v52 = int32(-1)
			} else {
				v52 = base.B2i32(base.Ui32(v47) < base.Ui32(v38))
			}
		}
	}
	return v52
}
func F_macaddr_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
	v6 = int32(8)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v9 = int32(16)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
	if v5<<(uint(v6)%32)|v8<<(uint(v9)%32)|v12 != v15<<(uint(v6)%32)|v18<<(uint(v9)%32)|v22 {
		v47 = v2
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+5)))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
		v27 = int32(8)
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+3)))
		v30 = int32(16)
		v33 = v25 | (v26<<(uint(v27)%32) | v29<<(uint(v30)%32))
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)))
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)))
		v42 = v34 | (v35<<(uint(v27)%32) | v38<<(uint(v30)%32))
		if base.Ui32(v33) < base.Ui32(v42) {
			v47 = v2
		} else {
			v47 = base.B2i32(base.Ui32(v33) <= base.Ui32(v42))
		}
	}
	return v47
}
func F_macaddr_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
	v7 = int32(8)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v10 = int32(16)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	v14 = v6<<(uint(v7)%32) | v9<<(uint(v10)%32) | v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)))
	v24 = v16<<(uint(v7)%32) | v19<<(uint(v10)%32) | v23
	if base.Ui32(v14) < base.Ui32(v24) {
		return int32(0)
	} else {
		if base.Ui32(v24) < base.Ui32(v14) {
			return int32(1)
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
			v33 = int32(8)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
			v36 = int32(16)
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)))
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)))
			return base.B2i32(base.Ui32(v40|(v41<<(uint(v33)%32)|v44<<(uint(v36)%32))) <= base.Ui32(v31|(v32<<(uint(v33)%32)|v35<<(uint(v36)%32))))
		}
	}
}
func F_macaddr_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
	v7 = int32(8)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v10 = int32(16)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	v14 = v6<<(uint(v7)%32) | v9<<(uint(v10)%32) | v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)))
	v24 = v16<<(uint(v7)%32) | v19<<(uint(v10)%32) | v23
	if base.Ui32(v14) < base.Ui32(v24) {
		return int32(1)
	} else {
		if base.Ui32(v24) < base.Ui32(v14) {
			return int32(0)
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
			v33 = int32(8)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
			v36 = int32(16)
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)))
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)))
			return base.B2i32(base.Ui32(v31|(v32<<(uint(v33)%32)|v35<<(uint(v36)%32))) < base.Ui32(v40|(v41<<(uint(v33)%32)|v44<<(uint(v36)%32))))
		}
	}
}
func F_macaddr_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v9 = m.G0
	v10 = int32(32)
	v11 = v9 - v10
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_palloc(m, v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v19
		v33 = F_pg_snprintf(m, v15, int32(32), int32(30032), v11)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v11 + int32(32)
			return v15
		}
	}
}
