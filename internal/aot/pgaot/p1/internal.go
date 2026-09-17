package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_internal_size(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v103 int64
	_ = v103
	var v115 int32
	_ = v115
	v4 = int32(0)
	if l1 <= v4 {
		v103 = int64(0)
	} else {
		v11 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+4)))
		v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0))))
		v15 = v11 - v12 + int64(1)
		if base.Ui32(l1) < base.Ui32(int32(3)) {
			v103 = v15
		} else {
			v21 = int32(base.Ui32(l1-int32(3)) >> (uint(int32(1)) % 32))
			if v21 == int32(0) {
				v80 = int32(2)
				v81 = v15
				v89 = l0 + v80<<(uint(int32(2))%32)
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(4))))
				if v90 == v93 {
					v103 = v81
				} else {
					v95 = int64(*(*int32)(unsafe.Add(mBase, uint32(v89)+4)))
					v103 = v81 + v95 - base.I64_extend_i32_s(v90) + int64(1)
				}
			} else {
				v25 = int32(1)
				v26 = v21 + v25
				v33 = int32(2)
				v34 = v15
				v39 = v4
				for {
					v42 = l0 + v33<<(uint(int32(2))%32)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v42-int32(4))))
					if v43 != v46 {
						v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+4)))
						v54 = v34 + v48 - base.I64_extend_i32_s(v43) + int64(1)
					} else {
						v54 = v34
					}
					v55 = int32(2)
					v59 = l0 + (v33+v55)<<(uint(v55)%32)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v59-int32(4))))
					if v60 != v63 {
						v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v59)+4)))
						v71 = v54 + v65 - base.I64_extend_i32_s(v60) + int64(1)
					} else {
						v71 = v54
					}
					v73 = v33 + int32(4)
					v75 = v39 + int32(2)
					if v75 != v26&int32(-2) {
						v33 = v73
						v34 = v71
						v39 = v75
						continue
					} else {
						break
					}
					break
				}
				if v26&v25 == int32(0) {
					v103 = v71
				} else {
					v80 = v73
					v81 = v71
					v89 = l0 + v80<<(uint(int32(2))%32)
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(4))))
					if v90 == v93 {
						v103 = v81
					} else {
						v95 = int64(*(*int32)(unsafe.Add(mBase, uint32(v89)+4)))
						v103 = v81 + v95 - base.I64_extend_i32_s(v90) + int64(1)
					}
				}
			}
		}
	}
	if base.Ui64(v103-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		v115 = int32(-1)
	} else {
		v115 = base.I32_wrap_i64(v103)
	}
	return v115
}
