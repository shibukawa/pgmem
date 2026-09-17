package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UnregisterSnapshotFromOwner(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	if l0 != 0 {
		F_ResourceOwnerForget(m, l1, l0, int32(_a_F_UnregisterSnapshotFromOwner_0))
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			F_UnregisterSnapshotNoOwner(m, l0)
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_UnreservedPLKeywords_hash_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v3 = int32(0)
	if l1 == v3 {
		v85 = int32(1)
		v91 = int32(0)
	} else {
		v13 = int32(1)
		if l1 != v13 {
			v21 = l0
			v22 = int32(0)
			v23 = v13
			v24 = v3
			for {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v30 = int32(32)
				v31 = v29 | v30
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v34 = v32 | v30
				v35 = int32(_a_F_UnreservedPLKeywords_hash_func_0)
				v40 = v31 + (v34+v23*v35)*v35
				v41 = int32(257)
				v46 = (v24*v41+v34)*v41 + v31
				v47 = int32(2)
				v48 = v21 + v47
				v50 = v22 + v47
				if v50 != l1&int32(-2) {
					v21 = v48
					v22 = v50
					v23 = v40
					v24 = v46
					continue
				} else {
					break
				}
				break
			}
			if l1&int32(1) == int32(0) {
				v73 = v40
				v74 = v46
			} else {
				v54 = v48
				v56 = v40
				v57 = v46
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v64 = v62 | int32(32)
				v73 = v64 + v56*int32(_a_F_UnreservedPLKeywords_hash_func_0)
				v74 = v57*int32(257) + v64
			}
		} else {
			v54 = l0
			v56 = v13
			v57 = v3
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
			v64 = v62 | int32(32)
			v73 = v64 + v56*int32(_a_F_UnreservedPLKeywords_hash_func_0)
			v74 = v57*int32(257) + v64
		}
		v79 = int32(167)
		v80 = base.I32_rem_u_s(v73, v79)
		v82 = base.I32_rem_u_s(v74, v79)
		v85 = v80
		v91 = v82
	}
	v92 = int32(1)
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v85<<(uint(v92)%32))+uint32(_c_F_UnreservedPLKeywords_hash_func[0]))))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91<<(uint(v92)%32))+uint32(_c_F_UnreservedPLKeywords_hash_func[0]))))
	return v96 + v101
}
func F_UpdateDecodingStats(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	v13 = m.G0
	v15 = v13 - int32(144)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
	if int64(0) < v18 {
		v29 = F_errstart(m, int32(13), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			if v29 != 0 {
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
				v32 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
				v33 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
				v34 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
				*(*int64)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v40
				*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v37
				*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v36
				*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v34
				*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v32
				*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v31
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
				F_errmsg_internal(m, int32(_a_F_UpdateDecodingStats_0), v15)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_UpdateDecodingStats_1), int32(1972), int32(_a_F_UpdateDecodingStats_2))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v65 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v65
						v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v67
						v69 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v69
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v71
						v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v73
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v75
						v77 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v77
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v79
						v82 = v15 + int32(72)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
						v90 = base.I32_div_s(v85-v87, int32(288))
						v93 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v90), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+24))
							v97 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+24)) = v96 + v97
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v95)+32))
							v101 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+32)) = v100 + v101
							v104 = *(*int64)(unsafe.Add(mBase, uint32(v95)+40))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v82)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+40)) = v104 + v105
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v82)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+48)) = v108 + v109
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v95)+56))
							v113 = *(*int64)(unsafe.Add(mBase, uint32(v82)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+56)) = v112 + v113
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v95)+64))
							v117 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+64)) = v116 + v117
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v95)+72))
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v82)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+72)) = v120 + v121
							v124 = *(*int64)(unsafe.Add(mBase, uint32(v95)+80))
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v82)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+80)) = v124 + v125
							F_pgstat_unlock_entry(m, v93)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								v130 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v17)+216)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+208)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+200)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+192)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = v130
								m.G0 = v15 + int32(144)
								return
							}
						}
					}
				}
			} else {
				v65 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v65
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v67
				v69 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v69
				v71 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v71
				v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v73
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v75
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v77
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v79
				v82 = v15 + int32(72)
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
				v90 = base.I32_div_s(v85-v87, int32(288))
				v93 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v90), int32(0))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
					v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+24))
					v97 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
					*(*int64)(unsafe.Add(mBase, uint32(v95)+24)) = v96 + v97
					v100 = *(*int64)(unsafe.Add(mBase, uint32(v95)+32))
					v101 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v95)+32)) = v100 + v101
					v104 = *(*int64)(unsafe.Add(mBase, uint32(v95)+40))
					v105 = *(*int64)(unsafe.Add(mBase, uint32(v82)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v95)+40)) = v104 + v105
					v108 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v82)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v95)+48)) = v108 + v109
					v112 = *(*int64)(unsafe.Add(mBase, uint32(v95)+56))
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v82)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v95)+56)) = v112 + v113
					v116 = *(*int64)(unsafe.Add(mBase, uint32(v95)+64))
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v95)+64)) = v116 + v117
					v120 = *(*int64)(unsafe.Add(mBase, uint32(v95)+72))
					v121 = *(*int64)(unsafe.Add(mBase, uint32(v82)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v95)+72)) = v120 + v121
					v124 = *(*int64)(unsafe.Add(mBase, uint32(v95)+80))
					v125 = *(*int64)(unsafe.Add(mBase, uint32(v82)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v95)+80)) = v124 + v125
					F_pgstat_unlock_entry(m, v93)
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						v130 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v17)+216)) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v17)+208)) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v17)+200)) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v17)+192)) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = v130
						m.G0 = v15 + int32(144)
						return
					}
				}
			}
		}
	} else {
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
		if int64(0) < v21 {
			v29 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				if v29 != 0 {
					v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
					v32 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
					*(*int64)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v40
					*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v37
					*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v36
					*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v34
					*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v33
					*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v32
					*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
					F_errmsg_internal(m, int32(_a_F_UpdateDecodingStats_0), v15)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_UpdateDecodingStats_1), int32(1972), int32(_a_F_UpdateDecodingStats_2))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
							*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v65
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
							*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v67
							v69 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
							*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v69
							v71 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
							*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v71
							v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
							*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v73
							v75 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
							*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v75
							v77 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
							*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v77
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
							*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v79
							v82 = v15 + int32(72)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
							v90 = base.I32_div_s(v85-v87, int32(288))
							v93 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v90), int32(0))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+24))
								v97 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
								*(*int64)(unsafe.Add(mBase, uint32(v95)+24)) = v96 + v97
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v95)+32))
								v101 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v95)+32)) = v100 + v101
								v104 = *(*int64)(unsafe.Add(mBase, uint32(v95)+40))
								v105 = *(*int64)(unsafe.Add(mBase, uint32(v82)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v95)+40)) = v104 + v105
								v108 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
								v109 = *(*int64)(unsafe.Add(mBase, uint32(v82)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v95)+48)) = v108 + v109
								v112 = *(*int64)(unsafe.Add(mBase, uint32(v95)+56))
								v113 = *(*int64)(unsafe.Add(mBase, uint32(v82)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v95)+56)) = v112 + v113
								v116 = *(*int64)(unsafe.Add(mBase, uint32(v95)+64))
								v117 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v95)+64)) = v116 + v117
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v95)+72))
								v121 = *(*int64)(unsafe.Add(mBase, uint32(v82)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v95)+72)) = v120 + v121
								v124 = *(*int64)(unsafe.Add(mBase, uint32(v95)+80))
								v125 = *(*int64)(unsafe.Add(mBase, uint32(v82)+56))
								*(*int64)(unsafe.Add(mBase, uint32(v95)+80)) = v124 + v125
								F_pgstat_unlock_entry(m, v93)
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return
								} else {
									v130 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v17)+216)) = v130
									*(*int64)(unsafe.Add(mBase, uint32(v17)+208)) = v130
									*(*int64)(unsafe.Add(mBase, uint32(v17)+200)) = v130
									*(*int64)(unsafe.Add(mBase, uint32(v17)+192)) = v130
									*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v130
									*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = v130
									*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v130
									*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = v130
									m.G0 = v15 + int32(144)
									return
								}
							}
						}
					}
				} else {
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v65
					v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v67
					v69 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v69
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v71
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v73
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v75
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v77
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v79
					v82 = v15 + int32(72)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
					v90 = base.I32_div_s(v85-v87, int32(288))
					v93 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v90), int32(0))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+24))
						v97 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
						*(*int64)(unsafe.Add(mBase, uint32(v95)+24)) = v96 + v97
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v95)+32))
						v101 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v95)+32)) = v100 + v101
						v104 = *(*int64)(unsafe.Add(mBase, uint32(v95)+40))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v82)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v95)+40)) = v104 + v105
						v108 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v82)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v95)+48)) = v108 + v109
						v112 = *(*int64)(unsafe.Add(mBase, uint32(v95)+56))
						v113 = *(*int64)(unsafe.Add(mBase, uint32(v82)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v95)+56)) = v112 + v113
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v95)+64))
						v117 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v95)+64)) = v116 + v117
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v95)+72))
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v82)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v95)+72)) = v120 + v121
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v95)+80))
						v125 = *(*int64)(unsafe.Add(mBase, uint32(v82)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v95)+80)) = v124 + v125
						F_pgstat_unlock_entry(m, v93)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							v130 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v17)+216)) = v130
							*(*int64)(unsafe.Add(mBase, uint32(v17)+208)) = v130
							*(*int64)(unsafe.Add(mBase, uint32(v17)+200)) = v130
							*(*int64)(unsafe.Add(mBase, uint32(v17)+192)) = v130
							*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v130
							*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = v130
							*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v130
							*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = v130
							m.G0 = v15 + int32(144)
							return
						}
					}
				}
			}
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
			if v24 <= int64(0) {
				m.G0 = v15 + int32(144)
				return
			} else {
				v29 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					if v29 != 0 {
						v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
						v32 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
						v33 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
						v34 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
						*(*int64)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v40
						*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v37
						*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v36
						*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v35
						*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v34
						*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v33
						*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v32
						*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v31
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
						F_errmsg_internal(m, int32(_a_F_UpdateDecodingStats_0), v15)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_UpdateDecodingStats_1), int32(1972), int32(_a_F_UpdateDecodingStats_2))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
								*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v65
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
								*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v67
								v69 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
								*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v69
								v71 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
								*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v71
								v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
								*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v73
								v75 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
								*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v75
								v77 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
								*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v77
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
								*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v79
								v82 = v15 + int32(72)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
								v90 = base.I32_div_s(v85-v87, int32(288))
								v93 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v90), int32(0))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
									v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+24))
									v97 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
									*(*int64)(unsafe.Add(mBase, uint32(v95)+24)) = v96 + v97
									v100 = *(*int64)(unsafe.Add(mBase, uint32(v95)+32))
									v101 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v95)+32)) = v100 + v101
									v104 = *(*int64)(unsafe.Add(mBase, uint32(v95)+40))
									v105 = *(*int64)(unsafe.Add(mBase, uint32(v82)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v95)+40)) = v104 + v105
									v108 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
									v109 = *(*int64)(unsafe.Add(mBase, uint32(v82)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v95)+48)) = v108 + v109
									v112 = *(*int64)(unsafe.Add(mBase, uint32(v95)+56))
									v113 = *(*int64)(unsafe.Add(mBase, uint32(v82)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v95)+56)) = v112 + v113
									v116 = *(*int64)(unsafe.Add(mBase, uint32(v95)+64))
									v117 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v95)+64)) = v116 + v117
									v120 = *(*int64)(unsafe.Add(mBase, uint32(v95)+72))
									v121 = *(*int64)(unsafe.Add(mBase, uint32(v82)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v95)+72)) = v120 + v121
									v124 = *(*int64)(unsafe.Add(mBase, uint32(v95)+80))
									v125 = *(*int64)(unsafe.Add(mBase, uint32(v82)+56))
									*(*int64)(unsafe.Add(mBase, uint32(v95)+80)) = v124 + v125
									F_pgstat_unlock_entry(m, v93)
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return
									} else {
										v130 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v17)+216)) = v130
										*(*int64)(unsafe.Add(mBase, uint32(v17)+208)) = v130
										*(*int64)(unsafe.Add(mBase, uint32(v17)+200)) = v130
										*(*int64)(unsafe.Add(mBase, uint32(v17)+192)) = v130
										*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v130
										*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = v130
										*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v130
										*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = v130
										m.G0 = v15 + int32(144)
										return
									}
								}
							}
						}
					} else {
						v65 = *(*int64)(unsafe.Add(mBase, uint32(v17)+160))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v65
						v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+168))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v67
						v69 = *(*int64)(unsafe.Add(mBase, uint32(v17)+176))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v69
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v17)+184))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v71
						v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+192))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v73
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v75
						v77 = *(*int64)(unsafe.Add(mBase, uint32(v17)+208))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v77
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+216))
						*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v79
						v82 = v15 + int32(72)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v87 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
						v90 = base.I32_div_s(v85-v87, int32(288))
						v93 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v90), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+24))
							v97 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+24)) = v96 + v97
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v95)+32))
							v101 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+32)) = v100 + v101
							v104 = *(*int64)(unsafe.Add(mBase, uint32(v95)+40))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v82)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+40)) = v104 + v105
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v82)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+48)) = v108 + v109
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v95)+56))
							v113 = *(*int64)(unsafe.Add(mBase, uint32(v82)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+56)) = v112 + v113
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v95)+64))
							v117 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+64)) = v116 + v117
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v95)+72))
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v82)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+72)) = v120 + v121
							v124 = *(*int64)(unsafe.Add(mBase, uint32(v95)+80))
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v82)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v95)+80)) = v124 + v125
							F_pgstat_unlock_entry(m, v93)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								v130 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v17)+216)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+208)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+200)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+192)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v130
								*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = v130
								m.G0 = v15 + int32(144)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_UtilityContainsQuery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = l0
	goto L1
L1:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	switch v5 - int32(241) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L7
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v3 = v27
	goto L1
L4:
	;
	return v24
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v21 == int32(6) {
		v26 = v20
		goto L3
	} else {
		goto L13
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v17 == int32(6) {
		v26 = v16
		goto L3
	} else {
		goto L12
	}
L7:
	;
	if v5 != int32(201) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v13 != int32(6) {
		v24 = v12
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v26 = v12
	goto L3
L12:
	;
	v24 = v16
	goto L4
L13:
	;
	v24 = v20
	goto L4
}
func F_uint32in_subr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, _c_F_uint32in_subr[0])) = v5
	v20 = F_strtox_2(m, l0, v11+int32(44), v5, int64(4294967295))
	mBase = m.M
	v21 = base.I32_wrap_i64(v20)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_uint32in_subr[0]))
	if v23 != 0 {
		v27 = base.B2i32(v23 != int32(68))
	} else {
		v27 = int32(0)
	}
	if v27 == int32(0) {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
		if v30 != l0 {
			if v23 == int32(68) {
				v55 = int32(0)
				v56 = F_errsave_start(m, l3)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v56 == int32(0) {
						v123 = v55
						m.G0 = v11 + int32(48)
						return v123
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
							F_errmsg(m, int32(_a_F_uint32in_subr_0), v11+int32(16))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, l3, int32(_a_F_uint32in_subr_1), int32(924), int32(_a_F_uint32in_subr_2))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v123 = v55
									m.G0 = v11 + int32(48)
									return v123
								}
							}
						}
					}
				}
			} else {
				if l1 == int32(0) {
					v83 = v30
					for {
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
						if base.B2i32(base.Ui32(v85-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v85 == int32(32)) != 0 {
							v83 = v83 + int32(1)
							continue
						} else {
							break
						}
						break
					}
					if v85 == int32(0) {
						v123 = v21
						m.G0 = v11 + int32(48)
						return v123
					} else {
						v97 = int32(0)
						v98 = F_errsave_start(m, l3)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							if v98 == int32(0) {
								v123 = v97
								m.G0 = v11 + int32(48)
								return v123
							} else {
								F_errcode(m, int32(33685634))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l2
									F_errmsg(m, int32(_a_F_uint32in_subr_3), v11+int32(32))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, l3, int32(_a_F_uint32in_subr_1), int32(940), int32(_a_F_uint32in_subr_2))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											v123 = v97
											m.G0 = v11 + int32(48)
											return v123
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
					v123 = v21
					m.G0 = v11 + int32(48)
					return v123
				}
			}
		} else {
			v33 = int32(0)
			v34 = F_errsave_start(m, l3)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				if v34 == int32(0) {
					v123 = v33
					m.G0 = v11 + int32(48)
					return v123
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
						F_errmsg(m, int32(_a_F_uint32in_subr_3), v11)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l3, int32(_a_F_uint32in_subr_1), int32(918), int32(_a_F_uint32in_subr_2))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v123 = v33
								m.G0 = v11 + int32(48)
								return v123
							}
						}
					}
				}
			}
		}
	} else {
		v33 = int32(0)
		v34 = F_errsave_start(m, l3)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			if v34 == int32(0) {
				v123 = v33
				m.G0 = v11 + int32(48)
				return v123
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
					F_errmsg(m, int32(_a_F_uint32in_subr_3), v11)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errsave_finish(m, l3, int32(_a_F_uint32in_subr_1), int32(918), int32(_a_F_uint32in_subr_2))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v123 = v33
							m.G0 = v11 + int32(48)
							return v123
						}
					}
				}
			}
		}
	}
}
func F_umask(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_umask(m, l0)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v2) {
		*(*int32)(unsafe.Add(mBase, _c_F_umask[0])) = int32(0) - v2
		v10 = int32(-1)
	} else {
		v10 = v2
	}
	return v10
}
func F_unaccent_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v2
	if v17 <= v2 {
		v128 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v128
L2:
	;
	v27 = v18
	v30 = v17
	goto L3
L3:
	;
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v117 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L5:
	;
	v114 = v30 - v105
	if int32(0) < v114 {
		v27 = v27 + v105
		v30 = v114
		goto L3
	} else {
		goto L31
	}
L6:
	;
	v95 = F_pg_mblen_range(m, v27, v17+v18)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L23
	} else {
		goto L28
	}
L7:
	;
	v38 = int32(0)
	v41 = v38
	v44 = v38
	v45 = v38
	v48 = v16
	goto L8
L8:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v27))))
	v56 = v48 + v53*int32(12)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v57 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v58 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L10:
	;
	goto L9
L11:
	;
	v58 = v56
	goto L13
L12:
	;
	v58 = v45
	goto L13
L13:
	;
	v60 = v41 + int32(1)
	if v57 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v60
	goto L16
L15:
	;
	v61 = v44
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v62 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v60) < base.Ui32(v30) {
		v41 = v60
		v44 = v61
		v45 = v58
		v48 = v62
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v68 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v71 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	F_appendBinaryStringInfo(m, v14, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L23
	} else {
		goto L27
	}
L22:
	;
	F_initStringInfo(m, v14)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v27 == v18 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	F_appendBinaryStringInfo(m, v14, v18, v27-v18)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v105 = v61
	goto L5
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v97 == int32(0) {
		v105 = v95
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_appendBinaryStringInfo(m, v14, v27, v95)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v105 = v95
	goto L5
L31:
	;
	goto L4
L32:
	;
	v128 = int32(0)
	goto L1
L33:
	;
	goto L34
L34:
	;
	v122 = F_palloc0(m, int32(16))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L23
	} else {
		goto L35
	}
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v125 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v125)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v124
	v128 = v122
	goto L1
}
func F_union_tuples(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	v4 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_union_tuples[0]))
	v21 = F_AllocSetContextCreateInternal(m, v16, int32(_a_F_union_tuples_0), v4, int32(_a_F_union_tuples_1), int32(_a_F_union_tuples_2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = int32(_a_F_union_tuples_3)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_union_tuples[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_union_tuples[0])) = v21
	v28 = F_brin_deform_tuple(m, l0, l2, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_union_tuples[0])) = v24
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(0) < v34 {
		goto L32
	} else {
		goto L33
	}
L5:
	;
	F_MemoryContextDelete(m, v21)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L31
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v35 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v34 <= int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v40 = int32(24)
	v47 = int32(0)
	goto L9
L9:
	;
	v60 = v47 * int32(20)
	v61 = v28 + v40 + v60
	v62 = v60 + (l1 + v40)
	v64 = v47 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v64)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
	if v67 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L5
L11:
	;
	v157 = v47 + int32(1)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v157 < v159 {
		v47 = v157
		goto L9
	} else {
		goto L30
	}
L12:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v134 = F_index_getprocinfo(m, v129, base.I32_extend16_s(v47+int32(1)), int32(4))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L28
	}
L13:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+3)))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
	if v73 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v76 = int32(1)
	goto L16
L15:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
	v76 = v75
	goto L16
L16:
	;
	if base.B2i32(v70 == int32(0))&(v76&int32(1)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
	if v82 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)) = uint8(v125)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
	if v127 != 0 {
		goto L11
	} else {
		goto L27
	}
L20:
	;
	if v70 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v85 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+2)) = uint16(v85)
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66))))
	if v87 == int32(0) {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v98 = int32(0)
	goto L23
L23:
	;
	v108 = v98 << (uint(int32(2)) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108+v109)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v66+int32(8)))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+10)))
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v113)+8)))
	v116 = F_datumCopy(m, v111, v114, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L11
L25:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v118+v108))) = v116
	v122 = v98 + int32(1)
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66))))
	if base.Ui32(v122) < base.Ui32(v123) {
		v98 = v122
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L12
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+248))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137+v64)))
	v140 = F_FunctionCall3Coll(m, v134, v139, l0, v62, v61)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L11
L30:
	;
	goto L10
L31:
	;
	return
L32:
	;
	v181 = int32(24)
	v189 = v4
	goto L35
L33:
	;
	goto L34
L34:
	;
	v283 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v283)
	F_MemoryContextDelete(m, v21)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L45
	}
L35:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v189<<(uint(int32(2))%32))))
	v204 = v189 * int32(20)
	v205 = l1 + v181 + v204
	v206 = v204 + (v28 + v181)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)) = uint8(v207)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)) = uint8(v209)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
	if v211 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L34
L37:
	;
	v265 = v189 + int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	if v265 < v267 {
		v189 = v265
		goto L35
	} else {
		goto L44
	}
L38:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202))))
	if v212 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v220 = int32(0)
	goto L40
L40:
	;
	v233 = v220 << (uint(int32(2)) % 32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233+v234)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v202+int32(8)+v233)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+10)))
	v240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+8)))
	v241 = F_datumCopy(m, v236, v239, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L37
L42:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v243+v233))) = v241
	v247 = v220 + int32(1)
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202))))
	if base.Ui32(v247) < base.Ui32(v248) {
		v220 = v247
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L36
L45:
	;
	return
}
func F_unique_key_recheck(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L10
	} else {
		goto L58
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L54
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v16 != int32(442) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19&int32(28) != int32(4) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	switch v19 & int32(3) {
	case 0:
		v49 = int32(24)
		goto L7
	default:
		goto L9
	case 2:
		goto L8
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13+v49)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+220)) = uint16(v52)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+216)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v58 = F_table_slot_create(m, v56, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L15
	}
L8:
	;
	v49 = int32(28)
	goto L7
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_unique_key_recheck_0)
	F_errmsg(m, int32(_a_F_unique_key_recheck_8), v11+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_unique_key_recheck_2), int32(83), int32(_a_F_unique_key_recheck_0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+220)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+212)) = uint16(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+188))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+44))
	v67 = m.T0[v66].(func(*base.Module, int32) int32)(m, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+80)) = uint8(v69)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_unique_key_recheck[0]))
	if v72 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_unique_key_recheck[1])))
	if v74&int32(1) == int32(0) {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+188))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+56))
	v88 = m.T0[v87].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v67, v11+int32(208), int32(_a_F_unique_key_recheck_7), v58, v11+int32(80), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	m.G0 = v11 + int32(224)
	return int32(0)
L22:
	;
	if v88 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_ExecDropSingleTupleTableSlot(m, v58)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+188))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	m.T0[v101].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+188))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	m.T0[v96].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+24))
	v107 = F_index_open(m, v105, int32(3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L31
	}
L29:
	;
	v123 = v11 + int32(80)
	v125 = v11 + int32(48)
	F_FormIndexDatum(m, v109, v58, v120, v123, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L10
	} else {
		goto L40
	}
L30:
	;
	v113 = F_CreateExecutorState(m)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L35
	}
L31:
	;
	v109 = F_BuildIndexInfo(m, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+76))
	if v111 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+92))
	if v112 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v120 = v2
	goto L29
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+152))
	if v115 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v118 = v115
	goto L38
L37:
	;
	v116 = F_MakePerTupleExprContext(m, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v58
	v120 = v113
	goto L29
L39:
	;
	v118 = v116
	goto L38
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v109)+92))
	if v129 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v120 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v136 = F_index_insert(m, v107, v123, v125, v11+int32(216), v128, int32(3), int32(0), v109)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L10
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_check_exclusion_constraint(m, v128, v107, v109, v11+int32(208), v11+int32(80), v11+int32(48), v120, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L10
	} else {
		goto L47
	}
L45:
	;
	F_index_insert_cleanup(m, v107, v109)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	goto L41
L47:
	;
	goto L41
L48:
	;
	F_FreeExecutorState(m, v120)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L10
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_ExecDropSingleTupleTableSlot(m, v58)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	F_relation_close(m, v107, int32(3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	goto L21
L54:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_unique_key_recheck_0)
	F_errmsg(m, int32(_a_F_unique_key_recheck_1), v11)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_unique_key_recheck_2), int32(62), int32(_a_F_unique_key_recheck_0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_unique_key_recheck_0)
	F_errmsg(m, int32(_a_F_unique_key_recheck_3), v11+int32(32))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_unique_key_recheck_2), int32(69), int32(_a_F_unique_key_recheck_0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_unique_key_recheck_4), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_unique_key_recheck_5), int32(1218), int32(_a_F_unique_key_recheck_6))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_unistr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v323 int32
	_ = v323
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v404 int32
	_ = v404
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v560 int32
	_ = v560
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v670 int32
	_ = v670
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v872 int32
	_ = v872
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_pg_detoast_datum_packed(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = int32(1)
	v31 = v26 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v34 = v32 & v30
	if v32 == v30 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L242
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L238
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L234
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L230
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L226
	}
L8:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v1038 = v1036 + int32(4)
	v1039 = F_palloc(m, v1038)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L221
	}
L9:
	;
	if v34 != 0 {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	F_initStringInfo(m, v23+int32(80))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v66 = int32(base.Ui32(v60)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L12:
	;
	F_initStringInfo(m, v23+int32(80))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L21
	}
L13:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v34 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L16:
	;
	if v37 == int32(18) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v48 = int32(16)
	goto L19
L18:
	;
	v48 = int32(0)
	goto L19
L19:
	;
	v66 = v48
	goto L10
L20:
	;
	v51 = int32(1)
	v66 = int32(base.Ui32(v32)>>(uint(v51)%32)) - v51
	goto L10
L21:
	;
	v73 = int32(4)
	goto L9
L22:
	;
	if v66 <= int32(0) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v73 = v66
	goto L9
L24:
	;
	v76 = v31
	goto L26
L25:
	;
	v76 = v26 + int32(4)
	goto L26
L26:
	;
	v79 = v76
	v80 = v73
	v84 = int32(0)
	goto L27
L27:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79))))
	if v97 == int32(92) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if v1000 != 0 {
		goto L7
	} else {
		goto L220
	}
L29:
	;
	if int32(0) < v996 {
		v79 = v995
		v80 = v996
		v84 = v1000
		goto L27
	} else {
		goto L219
	}
L30:
	;
	v995 = v991
	v996 = v990
	v1000 = int32(0)
	goto L29
L31:
	;
	if v80 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v84 != 0 {
		goto L7
	} else {
		goto L217
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L212
	}
L35:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v102 == int32(92) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v84 != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(v80) < base.Ui32(int32(5)) {
		goto L34
	} else {
		goto L41
	}
L39:
	;
	F_appendStringInfoChar(m, v23+int32(80), int32(92))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v110 = int32(2)
	v990 = v80 - v110
	v991 = v79 + v110
	goto L30
L41:
	;
	v119 = int32(0)
	goto L42
L42:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+(v79+int32(1))))))
	v151 = base.B2i32(base.Ui32(v140-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v140|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L44
L43:
	;
	if v151 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v151 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v153 = v119 + int32(1)
	if v153 != int32(4) {
		v119 = v153
		goto L42
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L43
L48:
	;
	goto L47
L49:
	;
	if base.Ui32(v80) < base.Ui32(int32(8)) {
		goto L34
	} else {
		goto L98
	}
L50:
	;
	if v80 == int32(5) {
		goto L34
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v226 = int32(-48)
	if v102 == int32(117) {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	if v102 != int32(117) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v166 = int32(0)
	goto L55
L55:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+(v79+int32(2))))))
	v198 = base.B2i32(base.Ui32(v187-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v187|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L57
L56:
	;
	if v198 == int32(0) {
		goto L34
	} else {
		goto L62
	}
L57:
	;
	if v198 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v200 = v166 + int32(1)
	if v200 != int32(4) {
		v166 = v200
		goto L55
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L56
L61:
	;
	goto L60
L62:
	;
	goto L52
L63:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	if base.Ui32((v256-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v277 = v226
		goto L70
	} else {
		goto L71
	}
L64:
	;
	v232 = int32(2)
	goto L66
L65:
	;
	v232 = int32(1)
	goto L66
L66:
	;
	v233 = v79 + v232
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if base.Ui32((v234-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v255 = v226
		goto L63
	} else {
		goto L67
	}
L67:
	;
	if base.Ui32((v234-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v255 = int32(-87)
		goto L63
	} else {
		goto L68
	}
L68:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v234-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v255 = int32(-55)
	goto L63
L70:
	;
	v278 = int32(-48)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+2)))
	if base.Ui32((v280-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v301 = v278
		goto L76
	} else {
		goto L77
	}
L71:
	;
	if base.Ui32((v256-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v277 = int32(-87)
	goto L70
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(int32(5)) < base.Ui32((v256-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	v277 = int32(-55)
	goto L70
L76:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+3)))
	if base.Ui32((v302-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v323 = v278
		goto L80
	} else {
		goto L81
	}
L77:
	;
	if base.Ui32((v280-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v301 = int32(-87)
		goto L76
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(int32(5)) < base.Ui32((v280-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v301 = int32(-55)
	goto L76
L80:
	;
	v336 = v323 + v302 + ((v277+v256)<<(uint(int32(8))%32) + (v234+v255)<<(uint(int32(12))%32) + (v280+v301)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_unistr_0)) <= base.Ui32(v336-int32(1)) {
		goto L6
	} else {
		goto L86
	}
L81:
	;
	if base.Ui32((v302-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v323 = int32(-87)
	goto L80
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(5)) < base.Ui32((v302-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v323 = int32(-55)
	goto L80
L86:
	;
	v342 = v336 & int32(_a_F_unistr_1)
	if v84 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v356&int32(-1024) == int32(_a_F_unistr_2) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	if v342 != int32(_a_F_unistr_3) {
		goto L7
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v342 == int32(_a_F_unistr_3) {
		goto L7
	} else {
		goto L92
	}
L91:
	;
	v356 = v84<<(uint(int32(10))%32)&int32(_a_F_unistr_4) | v336&int32(1023) + int32(_a_F_unistr_5)
	goto L87
L92:
	;
	v356 = v336
	goto L87
L93:
	;
	v371 = v356
	goto L95
L94:
	;
	v362 = v23 + int32(48)
	F_pg_unicode_to_server(m, v356, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v373 = v232 | int32(4)
	v995 = v373 + v79
	v996 = v80 - v373
	v1000 = v371
	goto L29
L96:
	;
	F_appendStringInfoString(m, v23+int32(80), v362)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v371 = int32(0)
	goto L95
L98:
	;
	if v102 != int32(43) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if base.B2i32(v102 != int32(85))|base.B2i32(base.Ui32(v80) < base.Ui32(int32(10))) != 0 {
		goto L34
	} else {
		goto L151
	}
L100:
	;
	v381 = v79 + int32(2)
	v383 = int32(0)
	goto L101
L101:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+v381))))
	v415 = base.B2i32(base.Ui32(v404-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v404|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L103
L102:
	;
	if v415 == int32(0) {
		goto L99
	} else {
		goto L108
	}
L103:
	;
	if v415 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v417 = v383 + int32(1)
	if v417 != int32(6) {
		v383 = v417
		goto L101
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	goto L102
L107:
	;
	goto L106
L108:
	;
	v423 = int32(-48)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	if base.Ui32((v425-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v446 = v423
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+3)))
	if base.Ui32((v447-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v468 = v423
		goto L113
	} else {
		goto L114
	}
L110:
	;
	if base.Ui32((v425-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v446 = int32(-87)
		goto L109
	} else {
		goto L111
	}
L111:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v425-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	v446 = int32(-55)
	goto L109
L113:
	;
	v469 = int32(-48)
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)))
	if base.Ui32((v471-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v492 = v469
		goto L119
	} else {
		goto L120
	}
L114:
	;
	if base.Ui32((v447-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v468 = int32(-87)
	goto L113
L116:
	;
	goto L117
L117:
	;
	if base.Ui32(int32(5)) < base.Ui32((v447-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	v468 = int32(-55)
	goto L113
L119:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)))
	if base.Ui32((v493-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v514 = v469
		goto L123
	} else {
		goto L124
	}
L120:
	;
	if base.Ui32((v471-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v492 = int32(-87)
		goto L119
	} else {
		goto L121
	}
L121:
	;
	if base.Ui32(int32(5)) < base.Ui32((v471-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	v492 = int32(-55)
	goto L119
L123:
	;
	v515 = int32(-48)
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	if base.Ui32((v517-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v538 = v515
		goto L129
	} else {
		goto L130
	}
L124:
	;
	if base.Ui32((v493-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v514 = int32(-87)
	goto L123
L126:
	;
	goto L127
L127:
	;
	if base.Ui32(int32(5)) < base.Ui32((v493-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v514 = int32(-55)
	goto L123
L129:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+7)))
	if base.Ui32((v539-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v560 = v515
		goto L133
	} else {
		goto L134
	}
L130:
	;
	if base.Ui32((v517-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v538 = int32(-87)
		goto L129
	} else {
		goto L131
	}
L131:
	;
	if base.Ui32(int32(5)) < base.Ui32((v517-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v538 = int32(-55)
	goto L129
L133:
	;
	v581 = v560 + v539 + ((v468+v447)<<(uint(int32(16))%32) + (v425+v446)<<(uint(int32(20))%32) + (v471+v492)<<(uint(int32(12))%32) + (v514+v493)<<(uint(int32(8))%32) + (v517+v538)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_unistr_0)) <= base.Ui32(v581-int32(1)) {
		goto L5
	} else {
		goto L139
	}
L134:
	;
	if base.Ui32((v539-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v560 = int32(-87)
	goto L133
L136:
	;
	goto L137
L137:
	;
	if base.Ui32(int32(5)) < base.Ui32((v539-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L138
	}
L138:
	;
	v560 = int32(-55)
	goto L133
L139:
	;
	v587 = v581 & int32(_a_F_unistr_1)
	if v84 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	if v601&int32(-1024) == int32(_a_F_unistr_2) {
		goto L146
	} else {
		goto L147
	}
L141:
	;
	if v587 != int32(_a_F_unistr_3) {
		goto L7
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if v587 == int32(_a_F_unistr_3) {
		goto L7
	} else {
		goto L145
	}
L144:
	;
	v601 = v84<<(uint(int32(10))%32)&int32(_a_F_unistr_4) | v581&int32(1023) + int32(_a_F_unistr_5)
	goto L140
L145:
	;
	v601 = v581
	goto L140
L146:
	;
	v616 = v601
	goto L148
L147:
	;
	v607 = v23 + int32(48)
	F_pg_unicode_to_server(m, v601, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v617 = int32(8)
	v995 = v79 + v617
	v996 = v80 - v617
	v1000 = v616
	goto L29
L149:
	;
	F_appendStringInfoString(m, v23+int32(80), v607)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v616 = int32(0)
	goto L148
L151:
	;
	v647 = v79 + int32(2)
	v649 = int32(0)
	goto L152
L152:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649+v647))))
	v681 = base.B2i32(base.Ui32(v670-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v670|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L154
L153:
	;
	if v681 == int32(0) {
		goto L34
	} else {
		goto L159
	}
L154:
	;
	if v681 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v683 = v649 + int32(1)
	if v683 != int32(8) {
		v649 = v683
		goto L152
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	goto L153
L158:
	;
	goto L157
L159:
	;
	v689 = int32(-48)
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647))))
	if base.Ui32((v691-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v712 = v689
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+3)))
	if base.Ui32((v713-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v734 = v689
		goto L164
	} else {
		goto L165
	}
L161:
	;
	if base.Ui32((v691-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v712 = int32(-87)
		goto L160
	} else {
		goto L162
	}
L162:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v691-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	v712 = int32(-55)
	goto L160
L164:
	;
	v735 = int32(-48)
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)))
	if base.Ui32((v737-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v758 = v735
		goto L170
	} else {
		goto L171
	}
L165:
	;
	if base.Ui32((v713-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v734 = int32(-87)
	goto L164
L167:
	;
	goto L168
L168:
	;
	if base.Ui32(int32(5)) < base.Ui32((v713-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	v734 = int32(-55)
	goto L164
L170:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)))
	if base.Ui32((v759-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v780 = v735
		goto L174
	} else {
		goto L175
	}
L171:
	;
	if base.Ui32((v737-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v758 = int32(-87)
		goto L170
	} else {
		goto L172
	}
L172:
	;
	if base.Ui32(int32(5)) < base.Ui32((v737-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	v758 = int32(-55)
	goto L170
L174:
	;
	v781 = int32(-48)
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	if base.Ui32((v783-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v804 = v781
		goto L180
	} else {
		goto L181
	}
L175:
	;
	if base.Ui32((v759-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v780 = int32(-87)
	goto L174
L177:
	;
	goto L178
L178:
	;
	if base.Ui32(int32(5)) < base.Ui32((v759-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L179
	}
L179:
	;
	v780 = int32(-55)
	goto L174
L180:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+7)))
	if base.Ui32((v805-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v826 = v781
		goto L184
	} else {
		goto L185
	}
L181:
	;
	if base.Ui32((v783-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v804 = int32(-87)
		goto L180
	} else {
		goto L182
	}
L182:
	;
	if base.Ui32(int32(5)) < base.Ui32((v783-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	v804 = int32(-55)
	goto L180
L184:
	;
	v827 = int32(-48)
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)))
	if base.Ui32((v829-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v850 = v827
		goto L190
	} else {
		goto L191
	}
L185:
	;
	if base.Ui32((v805-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v826 = int32(-87)
	goto L184
L187:
	;
	goto L188
L188:
	;
	if base.Ui32(int32(5)) < base.Ui32((v805-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	v826 = int32(-55)
	goto L184
L190:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+9)))
	if base.Ui32((v851-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v872 = v827
		goto L194
	} else {
		goto L195
	}
L191:
	;
	if base.Ui32((v829-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v850 = int32(-87)
		goto L190
	} else {
		goto L192
	}
L192:
	;
	if base.Ui32(int32(5)) < base.Ui32((v829-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L193
	}
L193:
	;
	v850 = int32(-55)
	goto L190
L194:
	;
	v901 = v872 + v851 + ((v734+v713)<<(uint(int32(24))%32) + (v691+v712)<<(uint(int32(28))%32) + (v737+v758)<<(uint(int32(20))%32) + (v780+v759)<<(uint(int32(16))%32) + (v783+v804)<<(uint(int32(12))%32) + (v826+v805)<<(uint(int32(8))%32) + (v829+v850)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_unistr_0)) <= base.Ui32(v901-int32(1)) {
		goto L4
	} else {
		goto L200
	}
L195:
	;
	if base.Ui32((v851-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v872 = int32(-87)
	goto L194
L197:
	;
	goto L198
L198:
	;
	if base.Ui32(int32(5)) < base.Ui32((v851-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	v872 = int32(-55)
	goto L194
L200:
	;
	v907 = v901 & int32(_a_F_unistr_1)
	if v84 != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	if v921&int32(-1024) == int32(_a_F_unistr_2) {
		goto L207
	} else {
		goto L208
	}
L202:
	;
	if v907 != int32(_a_F_unistr_3) {
		goto L7
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	if v907 == int32(_a_F_unistr_3) {
		goto L7
	} else {
		goto L206
	}
L205:
	;
	v921 = v84<<(uint(int32(10))%32)&int32(_a_F_unistr_4) | v901&int32(1023) + int32(_a_F_unistr_5)
	goto L201
L206:
	;
	v921 = v901
	goto L201
L207:
	;
	v936 = v921
	goto L209
L208:
	;
	v927 = v23 + int32(48)
	F_pg_unicode_to_server(m, v921, v927)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L210
	}
L209:
	;
	v937 = int32(10)
	v995 = v79 + v937
	v996 = v80 - v937
	v1000 = v936
	goto L29
L210:
	;
	F_appendStringInfoString(m, v23+int32(80), v927)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v936 = int32(0)
	goto L209
L212:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_unistr_6), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errhint(m, int32(_a_F_unistr_7), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_9), int32(_a_F_unistr_10))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	F_appendStringInfoChar(m, v23+int32(80), v97)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v985 = int32(1)
	v990 = v80 - v985
	v991 = v79 + v985
	goto L30
L219:
	;
	goto L28
L220:
	;
	goto L8
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1039))) = v1038 << (uint(int32(2)) % 32)
	if v1036 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	base.MemoryCopy(m, v1039+int32(4), v1035, v1036)
	goto L224
L223:
	;
	goto L224
L224:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	F_pfree(m, v1047)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	m.G0 = v23 + int32(96)
	return v1039
L226:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errmsg(m, int32(_a_F_unistr_11), int32(0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_12), int32(_a_F_unistr_10))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v336
	F_errmsg(m, int32(_a_F_unistr_13), v23)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_14), int32(_a_F_unistr_10))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v581
	F_errmsg(m, int32(_a_F_unistr_13), v23+int32(16))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_15), int32(_a_F_unistr_10))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v901
	F_errmsg(m, int32(_a_F_unistr_13), v23+int32(32))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_16), int32(_a_F_unistr_10))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	F_errmsg_internal(m, int32(_a_F_unistr_17), int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_18), int32(_a_F_unistr_19))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_updateClosestMatch(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(v6 == v3)|base.B2i32(l1 == v3) != 0 {
		return
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		if v12 == int32(0) {
			return
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v15 == int32(0) {
				return
			} else {
				v18 = F_strlen(m, v6)
				mBase = m.M
				if base.Ui32(int32(255)) < base.Ui32(v18) {
					return
				} else {
					v21 = F_strlen(m, l1)
					mBase = m.M
					if base.Ui32(int32(255)) < base.Ui32(v21) {
						return
					} else {
						v24 = int32(1)
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v29 = F_varstr_levenshtein_less_equal(m, v6, v18, l1, v21, v24, v24, v24, v27, v24)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v31 < v29 {
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v34 = F_strlen(m, v33)
								mBase = m.M
								if base.Ui32(int32(base.Ui32(v34)>>(uint(int32(1))%32))) < base.Ui32(v29) {
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if base.B2i32(v38 != int32(-1))&base.B2i32(v38 <= v29) != 0 {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v29
									}
								}
							}
							return
						}
					}
				}
			}
		}
	}
}
func F_update_controlfile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	v6 = m.G0
	v8 = v6 - int32(_a_F_update_controlfile_0)
	m.G0 = v8
	v10 = F_time(m)
	mBase = m.M
	v11 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+292)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v10
	v16 = m.Env.Pgmem_crc32c(m, v11, l1, int32(292))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+292)) = v16 ^ v11
	base.MemoryFill(m, v8+int32(1400), int32(0), int32(_a_F_update_controlfile_1))
	v26 = v8 + int32(1104)
	base.MemoryCopy(m, v26, l1, int32(296))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = int32(_a_F_update_controlfile_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l0
	v33 = v8 + int32(80)
	v38 = F_pg_snprintf(m, v33, int32(1024), int32(_a_F_update_controlfile_3), v8-int32(-64))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v41 = F_BasicOpenFile(m, v33, int32(2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L32
	}
L5:
	;
	if int32(0) <= v41 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[0])) = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(167772173)
	v52 = int32(_a_F_update_controlfile_4)
	v53 = F_write(m, v41, v26, v52)
	mBase = m.M
	if v53 != v52 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L28
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[0]))
	if v57 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v82 = int32(_a_F_update_controlfile_5)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[1]))
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v84
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = int32(167772171)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_controlfile[2])))
	if v92 != int32(1) {
		v106 = v84
		goto L20
	} else {
		goto L21
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[0])) = int32(51)
	goto L14
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v8 + int32(80)
	F_errmsg(m, int32(_a_F_update_controlfile_6), v8+int32(48))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_update_controlfile_7), int32(247), int32(_a_F_update_controlfile_8))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	if v106 != 0 {
		goto L4
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	goto L22
L22:
	;
	v97 = F_fsync(m, v41)
	mBase = m.M
	if v97 != int32(-1) {
		v106 = v97
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v106 = int32(-1)
	goto L20
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[0]))
	if v101 == int32(27) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = int32(0)
	v111 = F_close(m, v41)
	mBase = m.M
	if v111 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v8 + int32(_a_F_update_controlfile_0)
	return
L28:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(80)
	F_errmsg(m, int32(_a_F_update_controlfile_9), v8)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_update_controlfile_7), int32(226), int32(_a_F_update_controlfile_8))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v8 + int32(80)
	F_errmsg(m, int32(_a_F_update_controlfile_10), v8+int32(32))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_update_controlfile_7), int32(264), int32(_a_F_update_controlfile_8))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(80)
	F_errmsg(m, int32(_a_F_update_controlfile_11), v8+int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_update_controlfile_7), int32(278), int32(_a_F_update_controlfile_8))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_update_spins_per_delay(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_update_spins_per_delay[0]))
	v8 = base.I32_div_s(v3+l0*int32(15), int32(16))
	return v8
}
