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
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+68))
	if int32(0) < v6 {
		v10 = v6 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+68)) = v10
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v5+v10<<(uint(int32(2))%32))+72))
		v18 = v6*int32(24) + v5
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+80))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v19
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)+72))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v21
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v18-int32(-64))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v25
		return v15
	} else {
		v30 = F_core_yylex(m, l0, l0+int32(16), l1)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v37 = v35 + v36
			v38 = F_strlen(m, v37)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v38
			switch v30 - int32(265) {
			case 0:
				v42 = int32(265)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
				switch v44 - int32(35) {
				case 0:
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
					if v63 != 0 {
						v64 = int32(265)
					} else {
						v64 = int32(35)
					}
					return v64
				default:
					v71 = v42
					return v71
				case 25:
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
					if v47 != int32(60) {
						v71 = v42
						return v71
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+2)))
						if v50 != 0 {
							v71 = v42
							return v71
						} else {
							return int32(278)
						}
					}
				case 27:
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
					if v53 != int32(62) {
						v71 = v42
						return v71
					} else {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+2)))
						if v58 != 0 {
							v59 = int32(265)
						} else {
							v59 = int32(279)
						}
						return v59
					}
				}
			default:
				v71 = v30
				return v71
			case 2:
				v66 = F_pstrdup(m, v37)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v66
					v71 = int32(267)
					return v71
				}
			}
		}
	}
}
