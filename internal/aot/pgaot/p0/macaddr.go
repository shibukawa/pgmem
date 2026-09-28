package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr_fast_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v8 = base.I32_wrap_i64(l0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	v10 = int32(8)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v13 = int32(16)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)))
	v17 = v9<<(uint(v10)%32) | v12<<(uint(v13)%32) | v16
	v18 = base.I32_wrap_i64(l1)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+2)))
	v27 = v19<<(uint(v10)%32) | v22<<(uint(v13)%32) | v26
	if base.Ui32(v17) < base.Ui32(v27) {
		v54 = int32(-1)
	} else {
		if base.Ui32(v27) < base.Ui32(v17) {
			v54 = int32(1)
		} else {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
			v34 = int32(8)
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+3)))
			v37 = int32(16)
			v40 = v32 | (v33<<(uint(v34)%32) | v36<<(uint(v37)%32))
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+5)))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+3)))
			v49 = v41 | (v42<<(uint(v34)%32) | v45<<(uint(v37)%32))
			if base.Ui32(v40) < base.Ui32(v49) {
				v54 = int32(-1)
			} else {
				v54 = base.B2i32(base.Ui32(v49) < base.Ui32(v40))
			}
		}
	}
	return v54
}
func F_macaddr_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int64
	_ = v52
	v2 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v8 = int32(8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(16)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v15 = v7<<(uint(v8)%32) | v10<<(uint(v11)%32) | v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v25 = v17<<(uint(v8)%32) | v20<<(uint(v11)%32) | v24
	if base.Ui32(v15) < base.Ui32(v25) {
		v52 = v2
	} else {
		if base.Ui32(v25) < base.Ui32(v15) {
			v52 = int64(1)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
			v31 = int32(8)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
			v34 = int32(16)
			v37 = v29 | (v30<<(uint(v31)%32) | v33<<(uint(v34)%32))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
			v46 = v38 | (v39<<(uint(v31)%32) | v42<<(uint(v34)%32))
			if base.Ui32(v37) < base.Ui32(v46) {
				v52 = v2
			} else {
				v52 = base.I64_extend_i32_u(base.B2i32(base.Ui32(v46) < base.Ui32(v37)))
			}
		}
	}
	return v52
}
func F_macaddr_le(m *base.Module, l0 int32) int64 {
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
	var v53 int64
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v8 = int32(8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(16)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v15 = v7<<(uint(v8)%32) | v10<<(uint(v11)%32) | v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v25 = v17<<(uint(v8)%32) | v20<<(uint(v11)%32) | v24
	if base.Ui32(v15) < base.Ui32(v25) {
		v53 = int64(1)
	} else {
		if base.Ui32(v25) < base.Ui32(v15) {
			v53 = int64(0)
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
				v53 = int64(1)
			} else {
				v53 = base.I64_extend_i32_u(base.B2i32(base.Ui32(v38) <= base.Ui32(v47)))
			}
		}
	}
	return v53
}
func F_macaddr_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int64
	_ = v49
	v4 = int64(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
	v7 = int32(8)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v10 = int32(16)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)))
	if v6<<(uint(v7)%32)|v9<<(uint(v10)%32)|v13 != v16<<(uint(v7)%32)|v19<<(uint(v10)%32)|v23 {
		v49 = v4
	} else {
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)))
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
		v28 = int32(8)
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
		v31 = int32(16)
		v34 = v26 | (v27<<(uint(v28)%32) | v30<<(uint(v31)%32))
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)))
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)))
		v43 = v35 | (v36<<(uint(v28)%32) | v39<<(uint(v31)%32))
		if base.Ui32(v34) < base.Ui32(v43) {
			v49 = v4
		} else {
			v49 = base.I64_extend_i32_u(base.B2i32(base.Ui32(v43) < base.Ui32(v34)))
		}
	}
	return v49
}
func F_macaddr_or(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_palloc(m, int32(6))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		v13 = v11 | v12
		*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v13)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
		v17 = v15 | v16
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v17)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
		v21 = v19 | v20
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)) = uint8(v21)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+3)))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
		v25 = v23 | v24
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)) = uint8(v25)
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
		v29 = v27 | v28
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v29)
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+5)))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)))
		v33 = v31 | v32
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v33)
		return base.I64_extend_i32_u(v7)
	}
}
