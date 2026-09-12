package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_internal_yylex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
	if int32(0) < v6 {
		v10 = v6 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+72)) = v10
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v5+v10<<(uint(int32(2))%32))+76))
		v18 = v6*int32(24) + v5
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+84))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v19
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)+76))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v21
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v18)+68))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
		return v15
	} else {
		v28 = F_core_yylex(m, l0, l0+int32(16), l1)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v35 = v33 + v34
			v36 = F_strlen(m, v35)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v36
			switch v28 - int32(265) {
			case 0:
				v40 = int32(265)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
				switch v42 - int32(35) {
				case 0:
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
					if v61 != 0 {
						v62 = int32(265)
					} else {
						v62 = int32(35)
					}
					return v62
				default:
					v69 = v40
					return v69
				case 25:
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
					if v45 != int32(60) {
						v69 = v40
						return v69
					} else {
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
						if v48 != 0 {
							v69 = v40
							return v69
						} else {
							return int32(278)
						}
					}
				case 27:
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
					if v51 != int32(62) {
						v69 = v40
						return v69
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
						if v56 != 0 {
							v57 = int32(265)
						} else {
							v57 = int32(279)
						}
						return v57
					}
				}
			default:
				v69 = v28
				return v69
			case 2:
				v64 = F_pstrdup(m, v35)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64
					v69 = int32(267)
					return v69
				}
			}
		}
	}
}
