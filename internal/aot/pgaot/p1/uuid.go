package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_generate_random(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = m.Env.Pgmem_random_bytes(m, l0, int32(16))
	mBase = m.M
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v8 = v4&int32(15) | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v14 = v10&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v14)
	return
}
func F_uuid_generate_v1mc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	v3 = m.G0
	v5 = v3 - int32(128)
	m.G0 = v5
	F_uuid_generate_random(m, v5)
	mBase = m.M
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+10)))
	v10 = v8 | int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+10)) = uint8(v10)
	F_uuid_unparse(m, v5, v5+int32(16))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = v5 - int32(-64)
	F_uuid_generate_time(m, v19)
	mBase = m.M
	v22 = v5 + int32(80)
	F_uuid_unparse(m, v19, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = v5 + int32(103)
	v28 = v5 + int32(40)
	if (v28^v26)&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v105 = F_DirectFunctionCall1Coll(m, int32(3376), int32(0), v22)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L25
	}
L5:
	;
	goto L4
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v82)
	if v82&int32(255) == int32(0) {
		goto L5
	} else {
		goto L21
	}
L7:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v81 = v28
	v82 = v34
	v83 = v26
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v28&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = v28
	v40 = v26
	goto L13
L11:
	;
	v52 = v28
	v54 = v26
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v59 = int32(-2139062144)
	if (int32(16843008)-v56|v56)&v59 != v59 {
		v81 = v52
		v82 = v56
		v83 = v54
		goto L6
	} else {
		goto L17
	}
L13:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v41)
	if v41 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v52 = v48
	v54 = v46
	goto L12
L15:
	;
	v45 = int32(1)
	v46 = v40 + v45
	v48 = v38 + v45
	if v48&int32(3) != 0 {
		v38 = v48
		v40 = v46
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v64 = v52
	v65 = v56
	v66 = v54
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v65
	v68 = int32(4)
	v69 = v66 + v68
	v71 = v64 + v68
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v76 = int32(-2139062144)
	if (int32(16843008)-v73|v73)&v76 == v76 {
		v64 = v71
		v65 = v73
		v66 = v69
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v81 = v71
	v82 = v73
	v83 = v69
	goto L6
L20:
	;
	goto L19
L21:
	;
	v90 = v81
	v92 = v83
	goto L22
L22:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)) = uint8(v93)
	v95 = int32(1)
	if v93 != 0 {
		v90 = v90 + v95
		v92 = v92 + v95
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L5
L24:
	;
	goto L23
L25:
	;
	m.G0 = v5 + int32(128)
	return v105
}
func F_uuid_generate_v3(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14017(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_uuid_generate_v5(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14017(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_uuid_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v78 = v43<<(uint(v7)%64) | v43&v9<<(uint(v11)%64) | (v43&v14<<(uint(v16)%64) | v43&v18<<(uint(v20)%64)) | (int64(base.Ui64(v43)>>(uint(v20)%64))&v18 | int64(base.Ui64(v43)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v43)>>(uint(v11)%64))&v9 | int64(base.Ui64(v43)>>(uint(v7)%64))))
	if v41 == v78 {
		v81 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v82 = int64(56)
		v84 = int64(65280)
		v86 = int64(40)
		v89 = int64(16711680)
		v91 = int64(24)
		v93 = int64(4278190080)
		v95 = int64(8)
		v116 = v81<<(uint(v82)%64) | v81&v84<<(uint(v86)%64) | (v81&v89<<(uint(v91)%64) | v81&v93<<(uint(v95)%64)) | (int64(base.Ui64(v81)>>(uint(v95)%64))&v93 | int64(base.Ui64(v81)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v81)>>(uint(v86)%64))&v84 | int64(base.Ui64(v81)>>(uint(v82)%64))))
		v117 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
		v152 = v117<<(uint(v82)%64) | v117&v84<<(uint(v86)%64) | (v117&v89<<(uint(v91)%64) | v117&v93<<(uint(v95)%64)) | (int64(base.Ui64(v117)>>(uint(v95)%64))&v93 | int64(base.Ui64(v117)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v117)>>(uint(v86)%64))&v84 | int64(base.Ui64(v117)>>(uint(v82)%64))))
		if v116 == v152 {
			v162 = int32(0)
		} else {
			v154 = v152
			v155 = v116
			if base.Ui64(v155) < base.Ui64(v154) {
				v159 = int32(-1)
			} else {
				v159 = int32(1)
			}
			v162 = v159
		}
	} else {
		v154 = v78
		v155 = v41
		if base.Ui64(v155) < base.Ui64(v154) {
			v159 = int32(-1)
		} else {
			v159 = int32(1)
		}
		v162 = v159
	}
	return base.B2i32(int32(0) < v162)
}
func F_uuid_ns_dns(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14018(m, l0, int32(_a_F_uuid_ns_dns_0), int32(_a_F_uuid_ns_dns_1), int32(_a_F_uuid_ns_dns_2), int32(_a_F_uuid_ns_dns_3), int32(_a_F_uuid_ns_dns_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_uuid_ns_url(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14018(m, l0, int32(_a_F_uuid_ns_url_0), int32(_a_F_uuid_ns_url_1), int32(_a_F_uuid_ns_url_2), int32(_a_F_uuid_ns_url_3), int32(_a_F_uuid_ns_url_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
