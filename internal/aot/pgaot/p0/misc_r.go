package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReleaseAuxProcessResourcesCallback(m *base.Module, l0 int32, l1 int32) {
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResourcesCallback[0]))
	v5 = int32(1)
	v7 = base.B2i32(l0 == int32(0))
	F_ResourceOwnerReleaseInternal(m, v4, v5, v7, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResourcesCallback[0]))
		F_ResourceOwnerReleaseInternal(m, v12, int32(2), v7, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResourcesCallback[0]))
			F_ResourceOwnerReleaseInternal(m, v18, int32(3), v7, int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResourcesCallback[0]))
				v25 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)) = uint16(v25)
				return
			}
		}
	}
}
func F_ReleaseDeletionLock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(1259) {
		F_UnlockRelationOid(m, v3, int32(8))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		F_UnlockDatabaseObject(m, v4, v3, int32(8))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ReleaseLockIfHeld(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseLockIfHeld[0]))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(0)
	goto L3
L2:
	;
	v14 = v13
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = v16
	goto L5
L4:
	;
	return
L5:
	;
	v28 = v19 - int32(1)
	if v28 < int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
	if v37 < v36 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v33 = v15 + v28<<(uint(int32(4))%32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 != v14 {
		v19 = v28
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v40 = v16 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v36 - v37
	if v14 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v56 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = F_LockRelease(m, l0, v60, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L18
	}
L12:
	;
	F_ResourceOwnerForgetLock(m, v13, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v47 = v40
	goto L14
L14:
	;
	if v47 <= v28 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v47 = v46
	goto L14
L17:
	;
	v51 = v15 + v47<<(uint(int32(4))%32)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v54
	return
L18:
	;
	if v61 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v65 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	if v65 == int32(0) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errmsg_internal(m, int32(_a_F_ReleaseLockIfHeld_0), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ReleaseLockIfHeld_1), int32(2656), int32(_a_F_ReleaseLockIfHeld_2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L4
}
func F_RemoveLogrotateSignalFiles(m *base.Module) {
	var v2 int32
	_ = v2
	v2 = F_unlink(m, int32(_a_F_RemoveLogrotateSignalFiles_0))
	return
}
func F_RemoveNonParentXlogFiles(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v293 int32
	_ = v293
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	v21 = int64(*(*int32)(unsafe.Add(mBase, _c_F_RemoveNonParentXlogFiles[0])))
	v22 = base.I64_div_u_s(l0-int64(1), v21)
	v24 = base.I64_div_u_s(int64(4294967296), v21)
	v25 = base.I64_div_u_s(v22, v24)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+20)) = uint32(v25)
	v28 = v22 - v24*v25
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v28)
	v30 = base.I64_div_u_s(l0, v21)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v30
	v33 = v15 + int32(48)
	v38 = F_pg_snprintf(m, v33, int32(64), int32(_a_F_RemoveNonParentXlogFiles_0), v15+int32(16))
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
	v42 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v42 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v33
	F_errmsg_internal(m, int32(_a_F_RemoveNonParentXlogFiles_1), v15)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v54 = F_AllocateDir(m, int32(_a_F_RemoveNonParentXlogFiles_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_errfinish(m, int32(_a_F_RemoveNonParentXlogFiles_3), int32(3961), int32(_a_F_RemoveNonParentXlogFiles_4))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v57 = F_ReadDir(m, v54, int32(_a_F_RemoveNonParentXlogFiles_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v57 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = v15 + int32(48) | int32(8)
	v70 = v57
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_FreeDir(m, v54)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L66
	}
L14:
	;
	v78 = v70 + int32(19)
	v79 = F_strlen(m, v78)
	mBase = m.M
	if v79 != int32(24) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v278 = F_ReadDir(m, v54, int32(_a_F_RemoveNonParentXlogFiles_2))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L64
	}
L17:
	;
	v82 = int32(_a_F_RemoveNonParentXlogFiles_5)
	v86 = m.G0
	v88 = v86 - int32(32)
	v89 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+24)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v88))) = v89
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveNonParentXlogFiles[1])))
	if v97 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v165 != int32(24) {
		goto L16
	} else {
		goto L37
	}
L19:
	;
	v165 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveNonParentXlogFiles[2])))
	if v101 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v105 = v78
	goto L25
L23:
	;
	goto L24
L24:
	;
	v115 = v82
	v116 = v97
	goto L28
L25:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v111 == v97 {
		v105 = v105 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v165 = v105 - v78
	goto L18
L27:
	;
	goto L26
L28:
	;
	v123 = v88 + int32(base.Ui32(v116)>>(uint(int32(3))%32))&int32(28)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v125 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v124 | v125<<(uint(v116)%32)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v129 != 0 {
		v115 = v115 + v125
		v116 = v129
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v132 == int32(0) {
		v155 = v78
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v165 = v155 - v78
	goto L18
L32:
	;
	v136 = v78
	v137 = v132
	goto L33
L33:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v88+int32(base.Ui32(v137)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v145)>>(uint(v137)%32))&int32(1) == int32(0) {
		v155 = v136
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v155 = v153
	goto L31
L35:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v153 = v136 + int32(1)
	if v151 != 0 {
		v136 = v153
		v137 = v151
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v169 = v15 + int32(48)
	goto L40
L38:
	;
	if int32(0) <= v207-v208 {
		goto L16
	} else {
		goto L51
	}
L40:
	;
	goto L41
L41:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v176 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v177 = v78
	v178 = v169
	v179 = int32(8)
	v180 = v176
	goto L46
L43:
	;
	v203 = v169
	v207 = int32(0)
	goto L44
L44:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	goto L38
L45:
	;
	v203 = v198
	v207 = v200
	goto L44
L46:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if base.B2i32(v180 != v182)|base.B2i32(v182 == int32(0)) != 0 {
		v198 = v178
		v200 = v180
		goto L45
	} else {
		goto L48
	}
L47:
	;
	v198 = v192
	v200 = int32(0)
	goto L45
L48:
	;
	v188 = v179 - int32(1)
	if v188 == int32(0) {
		v198 = v178
		v200 = v180
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v191 = int32(1)
	v192 = v178 + v191
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	if v193 != 0 {
		v177 = v177 + v191
		v178 = v192
		v179 = v188
		v180 = v193
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v219 = v70 + int32(27)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if base.B2i32(v222 == int32(0))|base.B2i32(v222 != v225) != 0 {
		v243 = v222
		v244 = v225
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v243-v244 <= int32(0) {
		goto L16
	} else {
		goto L59
	}
L53:
	;
	goto L52
L54:
	;
	v228 = v219
	v229 = v64
	goto L55
L55:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	if v233 == int32(0) {
		v243 = v233
		v244 = v232
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v243 = v233
	v244 = v232
	goto L53
L57:
	;
	v236 = int32(1)
	if v233 == v232 {
		v228 = v228 + v236
		v229 = v229 + v236
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v248 = m.G0
	v250 = v248 - int32(1136)
	m.G0 = v250
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = int32(_a_F_RemoveNonParentXlogFiles_6)
	v256 = v250 + int32(112)
	v259 = F_pg_snprintf(m, v256, int32(1024), int32(_a_F_RemoveNonParentXlogFiles_7), v250)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v265 = F___fstatat(m, int32(-100), v256, v250+int32(16), int32(0))
	mBase = m.M
	goto L61
L61:
	;
	m.G0 = v250 + int32(1136)
	if v265 == int32(0) {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	F_RemoveXlogFile(m, v70, v30+int64(10), v15+int32(40), l1)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L16
L64:
	;
	if v278 != 0 {
		v70 = v278
		goto L14
	} else {
		goto L65
	}
L65:
	;
	goto L15
L66:
	;
	m.G0 = v15 + int32(112)
	return
}
func F_ReqShutdownXLOG(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	*(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[1]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[2]))
	if v17 == v13 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[3]))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_pgmem_kill(m, v13, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v21 + int32(16)
	goto L1
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v27)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[4]))
	v35 = F_write(m, v31, v21+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[5]))
	if v39 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_ResolveCminCmaxDuringDecoding(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v120 int32
	_ = v120
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int64
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int64
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int64
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1264)
	m.G0 = v20
	if l0 == v7 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L13
	} else {
		goto L163
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L13
	} else {
		goto L159
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L13
	} else {
		goto L156
	}
L4:
	;
	m.G0 = v20 + int32(1264)
	return v694
L5:
	;
	v694 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v25
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+152)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v20)+144)) = v27
	v32 = v20 + int32(144)
	if l3 < v25 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+160)) = uint16(v63)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+156)) = v65
	v67 = int32(0)
	v69 = F_hash_search(m, l0, v32, v67, v67)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(140)))) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(136)))) = v61
	goto L8
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[0]))
	v54 = v41 + (l3^int32(-1))<<(uint(int32(6))%32)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[1]))
	v54 = v48 + l3<<(uint(int32(6))%32) + int32(-64)
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	if v69 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v78 = int32(1)
	if v75 <= int32(3591) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v671 = v69
	goto L17
L17:
	;
	if l4 != 0 {
		goto L152
	} else {
		goto L153
	}
L18:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[2]))
	v153 = F_AllocateDir(m, int32(_a_F_ResolveCminCmaxDuringDecoding_0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L43
	}
L19:
	;
	goto L18
L20:
	;
	v149 = int32(0)
	goto L19
L21:
	;
	if base.B2i32(base.Ui32(v75-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v75-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v149 = v78
		goto L19
	} else {
		goto L42
	}
L22:
	;
	if v75 <= int32(2670) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v75 <= int32(_a_F_ResolveCminCmaxDuringDecoding_1) {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	switch v75 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v149 = v78
		goto L19
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L20
	default:
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v90 = v75 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v90))|base.B2i32(int32(1)<<(uint(v90)%32)&int32(226492515) == int32(0)) != 0 {
		goto L21
	} else {
		goto L30
	}
L28:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v75-int32(2396)) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v149 = v78
	goto L19
L30:
	;
	v149 = v78
	goto L19
L31:
	;
	if base.Ui32(v75-int32(3592)) < base.Ui32(int32(2)) {
		v149 = v78
		goto L19
	} else {
		goto L40
	}
L32:
	;
	v103 = v75 - int32(_a_F_ResolveCminCmaxDuringDecoding_2)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v103))|base.B2i32(int32(1)<<(uint(v103)%32)&int32(963) == int32(0)) != 0 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	switch v75 - int32(_a_F_ResolveCminCmaxDuringDecoding_3) {
	case 0, 1, 2, 3, 4, 59, 60:
		v149 = v78
		goto L19
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L20
	default:
		goto L36
	}
L35:
	;
	v149 = v78
	goto L19
L36:
	;
	if base.Ui32(v75-int32(_a_F_ResolveCminCmaxDuringDecoding_4)) < base.Ui32(int32(3)) {
		v149 = v78
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v120 = v75 - int32(_a_F_ResolveCminCmaxDuringDecoding_5)
	if base.Ui32(int32(15)) < base.Ui32(v120) {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	if int32(1)<<(uint(v120)%32)&int32(_a_F_ResolveCminCmaxDuringDecoding_6) != 0 {
		v149 = v78
		goto L19
	} else {
		goto L39
	}
L39:
	;
	goto L20
L40:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v75-int32(4060)) {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	v149 = v78
	goto L19
L42:
	;
	goto L20
L43:
	;
	v156 = F_ReadDir(m, v153, int32(_a_F_ResolveCminCmaxDuringDecoding_0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	if v156 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v149 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v393 = v7
	goto L47
L47:
	;
	F_FreeDir(m, v153)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L13
	} else {
		goto L106
	}
L48:
	;
	v159 = int32(0)
	goto L50
L49:
	;
	v159 = v151
	goto L50
L50:
	;
	v166 = v156
	v173 = v7
	goto L51
L51:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+19)))
	if v181 != int32(46) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v393 = v377
	goto L47
L53:
	;
	v382 = F_ReadDir(m, v153, int32(_a_F_ResolveCminCmaxDuringDecoding_0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L13
	} else {
		goto L104
	}
L54:
	;
	v194 = v166 + int32(19)
	v195 = int32(_a_F_ResolveCminCmaxDuringDecoding_7)
	goto L61
L55:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+20)))
	if v184 == int32(0) {
		v377 = v173
		goto L53
	} else {
		goto L56
	}
L56:
	;
	if v184 != int32(46) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+21)))
	if v189 == int32(0) {
		v377 = v173
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	if v233-v234 != 0 {
		v377 = v173
		goto L53
	} else {
		goto L72
	}
L61:
	;
	goto L62
L62:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v202 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v203 = v194
	v204 = v195
	v205 = int32(4)
	v206 = v202
	goto L67
L64:
	;
	v229 = v195
	v233 = int32(0)
	goto L65
L65:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	goto L59
L66:
	;
	v229 = v224
	v233 = v226
	goto L65
L67:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if base.B2i32(v206 != v208)|base.B2i32(v208 == int32(0)) != 0 {
		v224 = v204
		v226 = v206
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v224 = v218
	v226 = int32(0)
	goto L66
L69:
	;
	v214 = v205 - int32(1)
	if v214 == int32(0) {
		v224 = v204
		v226 = v206
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v217 = int32(1)
	v218 = v204 + v217
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)))
	if v219 != 0 {
		v203 = v203 + v217
		v204 = v218
		v205 = v214
		v206 = v219
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(132)))) = v20 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)))) = v20 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+124)) = v20 + int32(168)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+120)) = v20 + int32(172)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v20 + int32(184)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v20 + int32(204)
	v263 = F_sscanf(m, v194, int32(_a_F_ResolveCminCmaxDuringDecoding_8), v20+int32(112))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	if v263 != int32(6) {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v20)+204))
	if v267 != v159 {
		v377 = v173
		goto L53
	} else {
		goto L75
	}
L75:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v20)+184))
	if v269 != v75 {
		v377 = v173
		goto L53
	} else {
		goto L76
	}
L76:
	;
	v271 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+172)))
	v272 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+168)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	v274 = F_TransactionIdDidCommit(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	if v274 == int32(0) {
		v377 = v173
		goto L53
	} else {
		goto L78
	}
L78:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v20)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v280
	v286 = F_bsearch(m, v20+int32(240), v279, v278, int32(4), int32(185))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	if v286 == int32(0) {
		v377 = v173
		goto L53
	} else {
		goto L80
	}
L80:
	;
	v291 = F_palloc(m, int32(1032))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v291))) = v271<<(uint(int64(32))%64) | v272
	v298 = v291 + int32(8)
	if (v194^v298)&int32(3) != 0 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v373 = F_lappend(m, v173, v291)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L13
	} else {
		goto L103
	}
L83:
	;
	goto L82
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v353))) = uint8(v352)
	if v352&int32(255) == int32(0) {
		goto L83
	} else {
		goto L99
	}
L85:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v351 = v194
	v352 = v304
	v353 = v298
	goto L84
L86:
	;
	goto L87
L87:
	;
	if v194&int32(3) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v308 = v194
	v310 = v298
	goto L91
L89:
	;
	v322 = v194
	v324 = v298
	goto L90
L90:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v329 = int32(-2139062144)
	if (int32(16843008)-v326|v326)&v329 != v329 {
		v351 = v322
		v352 = v326
		v353 = v324
		goto L84
	} else {
		goto L95
	}
L91:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v311)
	if v311 == int32(0) {
		goto L83
	} else {
		goto L93
	}
L92:
	;
	v322 = v318
	v324 = v316
	goto L90
L93:
	;
	v315 = int32(1)
	v316 = v310 + v315
	v318 = v308 + v315
	if v318&int32(3) != 0 {
		v308 = v318
		v310 = v316
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v334 = v322
	v335 = v326
	v336 = v324
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336))) = v335
	v338 = int32(4)
	v339 = v336 + v338
	v341 = v334 + v338
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	v346 = int32(-2139062144)
	if (int32(16843008)-v343|v343)&v346 == v346 {
		v334 = v341
		v335 = v343
		v336 = v339
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v351 = v341
	v352 = v343
	v353 = v339
	goto L84
L98:
	;
	goto L97
L99:
	;
	v360 = v351
	v362 = v353
	goto L100
L100:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)) = uint8(v363)
	v365 = int32(1)
	if v363 != 0 {
		v360 = v360 + v365
		v362 = v362 + v365
		goto L100
	} else {
		goto L102
	}
L101:
	;
	goto L83
L102:
	;
	goto L101
L103:
	;
	v377 = v373
	goto L53
L104:
	;
	if v382 != 0 {
		v166 = v382
		v173 = v377
		goto L51
	} else {
		goto L105
	}
L105:
	;
	goto L52
L106:
	;
	F_list_sort(m, v393, int32(1019))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L13
	} else {
		goto L107
	}
L107:
	;
	if v393 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v659 = int32(0)
	v664 = F_hash_search(m, l0, v20+int32(144), v659, v659)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L13
	} else {
		goto L150
	}
L109:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v408 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v412 = v20 + int32(234)
	v414 = v20 + int32(216)
	v416 = v20 + int32(196)
	v418 = v20 + int32(228)
	v433 = v7
	goto L111
L111:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v436+v433<<(uint(int32(2))%32))))
	v443 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L13
	} else {
		goto L113
	}
L112:
	;
	goto L108
L113:
	;
	if v443 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v440 + int32(8)
	F_errmsg_internal(m, int32(_a_F_ResolveCminCmaxDuringDecoding_9), v20+int32(80))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L13
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = int32(_a_F_ResolveCminCmaxDuringDecoding_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v440 + int32(8)
	v467 = v20 + int32(240)
	v471 = F_pg_sprintf(m, v467, int32(_a_F_ResolveCminCmaxDuringDecoding_10), v20-int32(-64))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L13
	} else {
		goto L119
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_12), int32(_a_F_ResolveCminCmaxDuringDecoding_13))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L13
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v474 = F_OpenTransientFile(m, v467, int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	if v474 < int32(0) {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v478 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+184)) = v478
	*(*int64)(unsafe.Add(mBase, uint32(v20)+192)) = v478
	v482 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = v482
	v484 = int32(_a_F_ResolveCminCmaxDuringDecoding_14)
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = int32(167772205)
	v491 = F_read(m, v474, v20+int32(204), int32(36))
	mBase = m.M
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v482
	if v482 <= v491 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v634 = F_CloseTransientFile(m, v474)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L13
	} else {
		goto L146
	}
L123:
	;
	v501 = v491
	goto L126
L124:
	;
	goto L125
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L13
	} else {
		goto L142
	}
L126:
	;
	if v501 != int32(36) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L125
L128:
	;
	if v501 == int32(0) {
		goto L122
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v20)+212))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v541
	v543 = *(*int64)(unsafe.Add(mBase, uint32(v20)+204))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+184)) = v543
	v545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v416)+4)) = uint16(v545)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v547
	v550 = v20 + int32(184)
	v551 = int32(0)
	v553 = F_hash_search(m, l0, v550, v551, v551)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L13
	} else {
		goto L137
	}
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L13
	} else {
		goto L132
	}
L132:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L13
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v20 + int32(240)
	F_errmsg(m, int32(_a_F_ResolveCminCmaxDuringDecoding_15), v20+int32(32))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L13
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_16), int32(_a_F_ResolveCminCmaxDuringDecoding_17))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L13
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	v578 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+184)) = v578
	*(*int64)(unsafe.Add(mBase, uint32(v20)+192)) = v578
	v582 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = v582
	v584 = int32(_a_F_ResolveCminCmaxDuringDecoding_14)
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v585))) = int32(167772205)
	v591 = F_read(m, v474, v20+int32(204), int32(36))
	mBase = m.M
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = v582
	if v582 <= v591 {
		v501 = v591
		goto L126
	} else {
		goto L141
	}
L137:
	;
	if v553 == int32(0) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v414)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v557
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v414)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+184)) = v559
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v412)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v416)+4)) = uint16(v561)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v563
	v568 = F_hash_search(m, l0, v550, int32(1), v20+int32(180))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+180)))
	if v570 != 0 {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v553)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v568)+20)) = v571
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v553)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v568)+24)) = v573
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v553)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v568)+28)) = v575
	goto L136
L141:
	;
	goto L127
L142:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v20 + int32(240)
	F_errmsg(m, int32(_a_F_ResolveCminCmaxDuringDecoding_18), v20+int32(16))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L13
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_19), int32(_a_F_ResolveCminCmaxDuringDecoding_17))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	if v634 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_pfree(m, v440)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	v639 = v433 + int32(1)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v639 < v640 {
		v433 = v639
		goto L111
	} else {
		goto L149
	}
L149:
	;
	goto L112
L150:
	;
	if v664 == int32(0) {
		v694 = v659
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v671 = v664
	goto L17
L152:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v671)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v685
	goto L154
L153:
	;
	goto L154
L154:
	;
	v687 = int32(1)
	if l5 == int32(0) {
		v694 = v687
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v671)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v690
	v694 = v687
	goto L4
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v194
	F_errmsg_internal(m, int32(_a_F_ResolveCminCmaxDuringDecoding_20), v20+int32(96))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_21), int32(_a_F_ResolveCminCmaxDuringDecoding_13))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v20 + int32(240)
	F_errmsg(m, int32(_a_F_ResolveCminCmaxDuringDecoding_22), v20)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L13
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_23), int32(_a_F_ResolveCminCmaxDuringDecoding_17))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L13
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v20 + int32(240)
	F_errmsg(m, int32(_a_F_ResolveCminCmaxDuringDecoding_24), v20+int32(48))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L13
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_25), int32(_a_F_ResolveCminCmaxDuringDecoding_17))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RestoreSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreSnapshot[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = int32(2)
	v19 = v17 << (uint(v18) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v20 << (uint(v18) % 32)
	v26 = F_MemoryContextAlloc(m, v16, v19+v22+int32(72))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v26)+64)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v14
		v33 = int32(1)
		v34 = v13 & v33
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+29)) = uint8(v34)
		v37 = v12 & v33
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+28)) = uint8(v37)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v20
		v40 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v40
		*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v40
		*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v40
		v49 = l0 + int32(24)
		if v17 == v40 {
		} else {
			v53 = v26 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v53
			if v19 == int32(0) {
			} else {
				base.MemoryCopy(m, v53, v49, v19)
			}
		}
		if v20 <= int32(0) {
		} else {
			v62 = v17 << (uint(int32(2)) % 32)
			v65 = v26 + v62 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v65
			if v22 == int32(0) {
			} else {
				base.MemoryCopy(m, v65, v49+v62, v22)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v26)+44)) = int64(0)
		v75 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+30)) = uint8(v75)
		return v26
	}
}
func F_RestoreUserContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != int32(-1) {
		F_AtEOXact_GUC(m, int32(0), v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_RestoreUserContext[0])) = v10
			*(*int32)(unsafe.Add(mBase, _c_F_RestoreUserContext[1])) = v9
			return
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_RestoreUserContext[0])) = v10
		*(*int32)(unsafe.Add(mBase, _c_F_RestoreUserContext[1])) = v9
		return
	}
}
func F_r_LONG_2(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_find_among_b(m, l0, int32(_a_F_r_LONG_2_0), int32(7))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v4 != int32(0))
	}
}
func F_r_VOWEL_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 < v10 {
		v13 = v10
	} else {
		v13 = v11
	}
	if v10 == v13 {
		v54 = int32(-1)
	} else {
		v25 = int32(1)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v10))))
		if int32(117) < v28 {
			v49 = v25
		} else {
			v30 = v28 - int32(97)
			if v30 < int32(0) {
				v49 = v25
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v30)>>(uint(int32(3))%32)))+uint32(_c_F_r_VOWEL_1[0]))))
				if int32(base.Ui32(v36)>>(uint(v30&int32(7))%32))&int32(1) == int32(0) {
					v49 = v25
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
					v49 = int32(0)
				}
			}
		}
		v54 = v49
	}
	return base.B2i32(v54 == int32(0))
}
func F_r_e_ending_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v2
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 <= v11 {
		v93 = v2
		return v93
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v9-int32(1)))))
		if v17 != int32(101) {
			v93 = v2
			return v93
		} else {
			v21 = v9 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v9 <= v24 {
				v93 = v2
				return v93
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v34 <= v35 {
					v75 = int32(-1)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v34-int32(1)))))
					if int32(232) < v50 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - int32(1)
						v72 = int32(0)
					} else {
						v52 = v50 - int32(97)
						if v52 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - int32(1)
							v72 = int32(0)
						} else {
							v55 = int32(1)
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v52)>>(uint(int32(3))%32)))+uint32(_c_F_r_e_ending_1[0]))))
							if int32(base.Ui32(v59)>>(uint(v52&int32(7))%32))&v55 != 0 {
								v72 = v55
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - int32(1)
								v72 = int32(0)
							}
						}
					}
					v75 = v72
				}
				if v75 != 0 {
					v93 = v2
					return v93
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v76 + (v21 - v26)
					v80 = F_slice_del(m, l0)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						if v80 < int32(0) {
							v93 = v80
							return v93
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = int32(1)
							v89 = F_r_undouble_1(m, l0)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								v93 = v89
								return v93
							}
						}
					}
				}
			}
		}
	}
}
func F_r_en_ending_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v5 < v7 {
		v144 = v2
		return v144
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v17 <= v18 {
			v58 = int32(-1)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v17-int32(1)))))
			if int32(232) < v33 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 - int32(1)
				v55 = int32(0)
			} else {
				v35 = v33 - int32(97)
				if v35 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 - int32(1)
					v55 = int32(0)
				} else {
					v38 = int32(1)
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32)))+uint32(_c_F_r_en_ending_1[0]))))
					if int32(base.Ui32(v42)>>(uint(v35&int32(7))%32))&v38 != 0 {
						v55 = v38
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 - int32(1)
						v55 = int32(0)
					}
				}
			}
			v58 = v55
		}
		if v58 != 0 {
			v144 = v2
			return v144
		} else {
			v59 = v5 - v9
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v61 = v59 + v60
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
			v63 = int32(3)
			v65 = int32(0)
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v61-v68 < v63 {
				v78 = v65
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v74 = F_memcmp(m, v71+v61-v63, int32(_a_F_r_en_ending_1_0), v63)
				mBase = m.M
				if v74 != 0 {
					v78 = v65
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61 - v63
					v78 = int32(1)
				}
			}
			if v78 != 0 {
				v144 = v2
				return v144
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79 + v59
				v82 = F_slice_del(m, l0)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					if v82 < int32(0) {
						v144 = v82
						return v144
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v90 = v88 - int32(1)
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v90 <= v91 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v90))))
							if v97&int32(224) != int32(96) {
								return int32(0)
							} else {
								if int32(1)<<(uint(v97)%32)&int32(_a_F_r_en_ending_1_1) == int32(0) {
									return int32(0)
								} else {
									v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v115 = F_find_among_b(m, l0, int32(_a_F_r_en_ending_1_2), int32(3))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										if v115 == int32(0) {
											return int32(0)
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v123 = v121 + (v88 - v112)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v123
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123
											v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v123 <= v127 {
												v144 = int32(0)
												return v144
											} else {
												v129 = int32(1)
												v130 = v123 - v129
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v130
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v130
												v134 = F_slice_del(m, l0)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v134 {
														v141 = v129
													} else {
														v141 = v134 >> (uint(int32(31)) % 32) & v134
													}
													v144 = v141
													return v144
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_r_remove_suffix_2(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13986(m, l0, int32(_a_F_r_remove_suffix_2_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_r_shortv_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 <= v20 {
		v129 = int32(-1)
		v136 = v129
	} else {
		v37 = int32(1)
		v38 = v6 - v37
		v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21+v38))))
		v42 = v40 & int32(255)
		if base.B2i32(v38 == v20)|base.B2i32(int32(0) <= v40) != 0 {
			v100 = v42
			v104 = v37
		} else {
			v49 = v42 & int32(63)
			v51 = v6 - int32(2)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v51))))
			v55 = v53 << (uint(int32(6)) % 32)
			if base.B2i32(v51 != v20)&base.B2i32(base.Ui32(v53) < base.Ui32(int32(192))) == int32(0) {
				v100 = v55&int32(1984) | v49
				v104 = int32(2)
			} else {
				v68 = v55&int32(4032) | v49
				v70 = v6 - int32(3)
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v70))))
				if base.B2i32(v70 != v20)&base.B2i32(base.Ui32(v72) < base.Ui32(int32(224))) == int32(0) {
					v100 = v72<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v68
					v104 = int32(3)
				} else {
					v90 = int32(4)
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+v21-v90))))
					v100 = v72<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v92&int32(7)<<(uint(int32(18))%32) | v68
					v104 = v90
				}
			}
		}
		if int32(121) < v100 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v104
			v129 = int32(0)
			v136 = v129
		} else {
			v106 = v100 - int32(89)
			if v106 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v104
				v129 = int32(0)
				v136 = v129
			} else {
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v106)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[0]))))
				if int32(base.Ui32(v112)>>(uint(v106&int32(7))%32))&int32(1) == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v104
					v129 = int32(0)
					v136 = v129
				} else {
					v136 = v104
				}
			}
		}
	}
	if v136 != 0 {
		v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v400 = v398 + (v6 - v5)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400
		v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v400 <= v415 {
			v524 = int32(-1)
			v531 = v524
		} else {
			v432 = int32(1)
			v433 = v400 - v432
			v435 = int32(*(*int8)(unsafe.Add(mBase, uint32(v416+v433))))
			v437 = v435 & int32(255)
			if base.B2i32(v433 == v415)|base.B2i32(int32(0) <= v435) != 0 {
				v495 = v437
				v499 = v432
			} else {
				v444 = v437 & int32(63)
				v446 = v400 - int32(2)
				v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v446))))
				v450 = v448 << (uint(int32(6)) % 32)
				if base.B2i32(v446 != v415)&base.B2i32(base.Ui32(v448) < base.Ui32(int32(192))) == int32(0) {
					v495 = v450&int32(1984) | v444
					v499 = int32(2)
				} else {
					v463 = v450&int32(4032) | v444
					v465 = v400 - int32(3)
					v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v465))))
					if base.B2i32(v465 != v415)&base.B2i32(base.Ui32(v467) < base.Ui32(int32(224))) == int32(0) {
						v495 = v467<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v463
						v499 = int32(3)
					} else {
						v485 = int32(4)
						v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v416-v485))))
						v495 = v467<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v487&int32(7)<<(uint(int32(18))%32) | v463
						v499 = v485
					}
				}
			}
			if int32(121) < v495 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
				v524 = int32(0)
				v531 = v524
			} else {
				v501 = v495 - int32(97)
				if v501 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
					v524 = int32(0)
					v531 = v524
				} else {
					v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v501)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
					if int32(base.Ui32(v507)>>(uint(v501&int32(7))%32))&int32(1) == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
						v524 = int32(0)
						v531 = v524
					} else {
						v531 = v499
					}
				}
			}
		}
		if v531 != 0 {
			v664 = v2
		} else {
			v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v544 <= v545 {
				v653 = int32(-1)
				v660 = v653
			} else {
				v562 = int32(1)
				v563 = v544 - v562
				v565 = int32(*(*int8)(unsafe.Add(mBase, uint32(v546+v563))))
				v567 = v565 & int32(255)
				if base.B2i32(v563 == v545)|base.B2i32(int32(0) <= v565) != 0 {
					v625 = v567
					v629 = v562
				} else {
					v574 = v567 & int32(63)
					v576 = v544 - int32(2)
					v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v576))))
					v580 = v578 << (uint(int32(6)) % 32)
					if base.B2i32(v576 != v545)&base.B2i32(base.Ui32(v578) < base.Ui32(int32(192))) == int32(0) {
						v625 = v580&int32(1984) | v574
						v629 = int32(2)
					} else {
						v593 = v580&int32(4032) | v574
						v595 = v544 - int32(3)
						v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v595))))
						if base.B2i32(v595 != v545)&base.B2i32(base.Ui32(v597) < base.Ui32(int32(224))) == int32(0) {
							v625 = v597<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v593
							v629 = int32(3)
						} else {
							v615 = int32(4)
							v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+v546-v615))))
							v625 = v597<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v617&int32(7)<<(uint(int32(18))%32) | v593
							v629 = v615
						}
					}
				}
				if int32(121) < v625 {
					v660 = v629
				} else {
					v631 = v625 - int32(97)
					if v631 < int32(0) {
						v660 = v629
					} else {
						v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v631)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
						if int32(base.Ui32(v637)>>(uint(v631&int32(7))%32))&int32(1) == int32(0) {
							v660 = v629
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v544 - v629
							v653 = int32(0)
							v660 = v653
						}
					}
				}
			}
			if v660 != 0 {
				v664 = v2
			} else {
				v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v664 = base.B2i32(v661 <= v662)
			}
		}
		return v664
	} else {
		v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v149 <= v150 {
			v258 = int32(-1)
			v265 = v258
		} else {
			v167 = int32(1)
			v168 = v149 - v167
			v170 = int32(*(*int8)(unsafe.Add(mBase, uint32(v151+v168))))
			v172 = v170 & int32(255)
			if base.B2i32(v168 == v150)|base.B2i32(int32(0) <= v170) != 0 {
				v230 = v172
				v234 = v167
			} else {
				v179 = v172 & int32(63)
				v181 = v149 - int32(2)
				v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v181))))
				v185 = v183 << (uint(int32(6)) % 32)
				if base.B2i32(v181 != v150)&base.B2i32(base.Ui32(v183) < base.Ui32(int32(192))) == int32(0) {
					v230 = v185&int32(1984) | v179
					v234 = int32(2)
				} else {
					v198 = v185&int32(4032) | v179
					v200 = v149 - int32(3)
					v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v200))))
					if base.B2i32(v200 != v150)&base.B2i32(base.Ui32(v202) < base.Ui32(int32(224))) == int32(0) {
						v230 = v202<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v198
						v234 = int32(3)
					} else {
						v220 = int32(4)
						v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v151-v220))))
						v230 = v202<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v222&int32(7)<<(uint(int32(18))%32) | v198
						v234 = v220
					}
				}
			}
			if int32(121) < v230 {
				v265 = v234
			} else {
				v236 = v230 - int32(97)
				if v236 < int32(0) {
					v265 = v234
				} else {
					v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v236)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
					if int32(base.Ui32(v242)>>(uint(v236&int32(7))%32))&int32(1) == int32(0) {
						v265 = v234
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v149 - v234
						v258 = int32(0)
						v265 = v258
					}
				}
			}
		}
		if v265 != 0 {
			v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v400 = v398 + (v6 - v5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400
			v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v400 <= v415 {
				v524 = int32(-1)
				v531 = v524
			} else {
				v432 = int32(1)
				v433 = v400 - v432
				v435 = int32(*(*int8)(unsafe.Add(mBase, uint32(v416+v433))))
				v437 = v435 & int32(255)
				if base.B2i32(v433 == v415)|base.B2i32(int32(0) <= v435) != 0 {
					v495 = v437
					v499 = v432
				} else {
					v444 = v437 & int32(63)
					v446 = v400 - int32(2)
					v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v446))))
					v450 = v448 << (uint(int32(6)) % 32)
					if base.B2i32(v446 != v415)&base.B2i32(base.Ui32(v448) < base.Ui32(int32(192))) == int32(0) {
						v495 = v450&int32(1984) | v444
						v499 = int32(2)
					} else {
						v463 = v450&int32(4032) | v444
						v465 = v400 - int32(3)
						v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v465))))
						if base.B2i32(v465 != v415)&base.B2i32(base.Ui32(v467) < base.Ui32(int32(224))) == int32(0) {
							v495 = v467<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v463
							v499 = int32(3)
						} else {
							v485 = int32(4)
							v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v416-v485))))
							v495 = v467<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v487&int32(7)<<(uint(int32(18))%32) | v463
							v499 = v485
						}
					}
				}
				if int32(121) < v495 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
					v524 = int32(0)
					v531 = v524
				} else {
					v501 = v495 - int32(97)
					if v501 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
						v524 = int32(0)
						v531 = v524
					} else {
						v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v501)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
						if int32(base.Ui32(v507)>>(uint(v501&int32(7))%32))&int32(1) == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
							v524 = int32(0)
							v531 = v524
						} else {
							v531 = v499
						}
					}
				}
			}
			if v531 != 0 {
				v664 = v2
			} else {
				v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v544 <= v545 {
					v653 = int32(-1)
					v660 = v653
				} else {
					v562 = int32(1)
					v563 = v544 - v562
					v565 = int32(*(*int8)(unsafe.Add(mBase, uint32(v546+v563))))
					v567 = v565 & int32(255)
					if base.B2i32(v563 == v545)|base.B2i32(int32(0) <= v565) != 0 {
						v625 = v567
						v629 = v562
					} else {
						v574 = v567 & int32(63)
						v576 = v544 - int32(2)
						v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v576))))
						v580 = v578 << (uint(int32(6)) % 32)
						if base.B2i32(v576 != v545)&base.B2i32(base.Ui32(v578) < base.Ui32(int32(192))) == int32(0) {
							v625 = v580&int32(1984) | v574
							v629 = int32(2)
						} else {
							v593 = v580&int32(4032) | v574
							v595 = v544 - int32(3)
							v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v595))))
							if base.B2i32(v595 != v545)&base.B2i32(base.Ui32(v597) < base.Ui32(int32(224))) == int32(0) {
								v625 = v597<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v593
								v629 = int32(3)
							} else {
								v615 = int32(4)
								v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+v546-v615))))
								v625 = v597<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v617&int32(7)<<(uint(int32(18))%32) | v593
								v629 = v615
							}
						}
					}
					if int32(121) < v625 {
						v660 = v629
					} else {
						v631 = v625 - int32(97)
						if v631 < int32(0) {
							v660 = v629
						} else {
							v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v631)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
							if int32(base.Ui32(v637)>>(uint(v631&int32(7))%32))&int32(1) == int32(0) {
								v660 = v629
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v544 - v629
								v653 = int32(0)
								v660 = v653
							}
						}
					}
				}
				if v660 != 0 {
					v664 = v2
				} else {
					v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v664 = base.B2i32(v661 <= v662)
				}
			}
			return v664
		} else {
			v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v278 <= v279 {
				v388 = int32(-1)
				v395 = v388
			} else {
				v296 = int32(1)
				v297 = v278 - v296
				v299 = int32(*(*int8)(unsafe.Add(mBase, uint32(v280+v297))))
				v301 = v299 & int32(255)
				if base.B2i32(v297 == v279)|base.B2i32(int32(0) <= v299) != 0 {
					v359 = v301
					v363 = v296
				} else {
					v308 = v301 & int32(63)
					v310 = v278 - int32(2)
					v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v310))))
					v314 = v312 << (uint(int32(6)) % 32)
					if base.B2i32(v310 != v279)&base.B2i32(base.Ui32(v312) < base.Ui32(int32(192))) == int32(0) {
						v359 = v314&int32(1984) | v308
						v363 = int32(2)
					} else {
						v327 = v314&int32(4032) | v308
						v329 = v278 - int32(3)
						v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v329))))
						if base.B2i32(v329 != v279)&base.B2i32(base.Ui32(v331) < base.Ui32(int32(224))) == int32(0) {
							v359 = v331<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v327
							v363 = int32(3)
						} else {
							v349 = int32(4)
							v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278+v280-v349))))
							v359 = v331<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v351&int32(7)<<(uint(int32(18))%32) | v327
							v363 = v349
						}
					}
				}
				if int32(121) < v359 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v278 - v363
					v388 = int32(0)
					v395 = v388
				} else {
					v365 = v359 - int32(97)
					if v365 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v278 - v363
						v388 = int32(0)
						v395 = v388
					} else {
						v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v365)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
						if int32(base.Ui32(v371)>>(uint(v365&int32(7))%32))&int32(1) == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v278 - v363
							v388 = int32(0)
							v395 = v388
						} else {
							v395 = v363
						}
					}
				}
			}
			if v395 != 0 {
				v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v400 = v398 + (v6 - v5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400
				v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v400 <= v415 {
					v524 = int32(-1)
					v531 = v524
				} else {
					v432 = int32(1)
					v433 = v400 - v432
					v435 = int32(*(*int8)(unsafe.Add(mBase, uint32(v416+v433))))
					v437 = v435 & int32(255)
					if base.B2i32(v433 == v415)|base.B2i32(int32(0) <= v435) != 0 {
						v495 = v437
						v499 = v432
					} else {
						v444 = v437 & int32(63)
						v446 = v400 - int32(2)
						v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v446))))
						v450 = v448 << (uint(int32(6)) % 32)
						if base.B2i32(v446 != v415)&base.B2i32(base.Ui32(v448) < base.Ui32(int32(192))) == int32(0) {
							v495 = v450&int32(1984) | v444
							v499 = int32(2)
						} else {
							v463 = v450&int32(4032) | v444
							v465 = v400 - int32(3)
							v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v465))))
							if base.B2i32(v465 != v415)&base.B2i32(base.Ui32(v467) < base.Ui32(int32(224))) == int32(0) {
								v495 = v467<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v463
								v499 = int32(3)
							} else {
								v485 = int32(4)
								v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v416-v485))))
								v495 = v467<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v487&int32(7)<<(uint(int32(18))%32) | v463
								v499 = v485
							}
						}
					}
					if int32(121) < v495 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
						v524 = int32(0)
						v531 = v524
					} else {
						v501 = v495 - int32(97)
						if v501 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
							v524 = int32(0)
							v531 = v524
						} else {
							v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v501)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
							if int32(base.Ui32(v507)>>(uint(v501&int32(7))%32))&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v499
								v524 = int32(0)
								v531 = v524
							} else {
								v531 = v499
							}
						}
					}
				}
				if v531 != 0 {
					v664 = v2
				} else {
					v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v544 <= v545 {
						v653 = int32(-1)
						v660 = v653
					} else {
						v562 = int32(1)
						v563 = v544 - v562
						v565 = int32(*(*int8)(unsafe.Add(mBase, uint32(v546+v563))))
						v567 = v565 & int32(255)
						if base.B2i32(v563 == v545)|base.B2i32(int32(0) <= v565) != 0 {
							v625 = v567
							v629 = v562
						} else {
							v574 = v567 & int32(63)
							v576 = v544 - int32(2)
							v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v576))))
							v580 = v578 << (uint(int32(6)) % 32)
							if base.B2i32(v576 != v545)&base.B2i32(base.Ui32(v578) < base.Ui32(int32(192))) == int32(0) {
								v625 = v580&int32(1984) | v574
								v629 = int32(2)
							} else {
								v593 = v580&int32(4032) | v574
								v595 = v544 - int32(3)
								v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v595))))
								if base.B2i32(v595 != v545)&base.B2i32(base.Ui32(v597) < base.Ui32(int32(224))) == int32(0) {
									v625 = v597<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v593
									v629 = int32(3)
								} else {
									v615 = int32(4)
									v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+v546-v615))))
									v625 = v597<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v617&int32(7)<<(uint(int32(18))%32) | v593
									v629 = v615
								}
							}
						}
						if int32(121) < v625 {
							v660 = v629
						} else {
							v631 = v625 - int32(97)
							if v631 < int32(0) {
								v660 = v629
							} else {
								v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v631)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
								if int32(base.Ui32(v637)>>(uint(v631&int32(7))%32))&int32(1) == int32(0) {
									v660 = v629
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v544 - v629
									v653 = int32(0)
									v660 = v653
								}
							}
						}
					}
					if v660 != 0 {
						v664 = v2
					} else {
						v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v664 = base.B2i32(v661 <= v662)
					}
				}
				return v664
			} else {
				return int32(1)
			}
		}
	}
}
func F_r_undouble_4(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L4
L1:
	;
	return v134
L2:
	;
	if v58 < int32(0) {
		v134 = v2
		goto L1
	} else {
		goto L21
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	v13 = v5
	v15 = int32(1)
	goto L9
L8:
	;
	v58 = v40
	goto L2
L9:
	;
	if v13 <= v6 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v58 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v20 = v13 - int32(1)
	v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4+v20))))
	if base.B2i32(int32(0) <= v22)|base.B2i32(v20 <= v6) != 0 {
		v40 = v20
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = int32(1)
	if v44 < v15 {
		v13 = v40
		v15 = v15 - v44
		goto L9
	} else {
		goto L20
	}
L15:
	;
	v28 = v20
	goto L16
L16:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+v28))))
	if base.Ui32(int32(191)) < base.Ui32(v33) {
		v40 = v28
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v40 = v6
	goto L14
L18:
	;
	v37 = v28 - int32(1)
	if v6 < v37 {
		v28 = v37
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L10
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v58
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L24
L22:
	;
	if v116 < int32(0) {
		v134 = v2
		goto L1
	} else {
		goto L41
	}
L24:
	;
	goto L25
L25:
	;
	goto L26
L26:
	;
	v71 = v58
	v73 = int32(1)
	goto L29
L28:
	;
	v116 = v98
	goto L22
L29:
	;
	if v71 <= v64 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v116 = int32(-1)
	goto L22
L32:
	;
	goto L33
L33:
	;
	v78 = v71 - int32(1)
	v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63+v78))))
	if base.B2i32(int32(0) <= v80)|base.B2i32(v78 <= v64) != 0 {
		v98 = v78
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v102 = int32(1)
	if v102 < v73 {
		v71 = v98
		v73 = v73 - v102
		goto L29
	} else {
		goto L40
	}
L35:
	;
	v86 = v78
	goto L36
L36:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v86))))
	if base.Ui32(int32(191)) < base.Ui32(v91) {
		v98 = v86
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v98 = v64
	goto L34
L38:
	;
	v95 = v86 - int32(1)
	if v64 < v95 {
		v86 = v95
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L30
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v122 = F_slice_del(m, l0)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return int32(0)
L43:
	;
	if int32(0) <= v122 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v131 = int32(1)
	goto L46
L45:
	;
	v131 = v122 >> (uint(int32(31)) % 32) & v122
	goto L46
L46:
	;
	v134 = v131
	goto L1
}
func F_rainbow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	if l2 != int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(24)
	v19 = v13 + v14*v15 + v15
	if base.Ui32(v19) <= base.Ui32(v13) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_rainbow[0]))
	if v125 != 0 {
		goto L40
	} else {
		goto L41
	}
L5:
	;
	v27 = v13
	v28 = int32(0)
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L1
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v33&int32(1)|v33&int32(2) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v122 = v27 + int32(24)
	if base.Ui32(v122) < base.Ui32(v19) {
		v27 = v122
		v28 = v28 + int32(1)
		goto L6
	} else {
		goto L39
	}
L10:
	;
	v39 = int32(_a_F_rainbow_0)
	v40 = v28 & v39
	if v40 == l2&v39 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
	if v44 == v40 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_rainbow[0]))
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v50 <= v51 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	F_createarc(m, l0, int32(112), base.I32_extend16_s(v28), l3, l4)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L16
	} else {
		goto L38
	}
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v53 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v74 == int32(0) {
		goto L18
	} else {
		goto L30
	}
L22:
	;
	v61 = v53
	goto L23
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if v66 != l4 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L18
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	if v73 != 0 {
		v61 = v73
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	if v68 != v40 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v70 == int32(112) {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	goto L24
L30:
	;
	v82 = v74
	goto L31
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	if v87 != l3 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L18
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v94 != 0 {
		v82 = v94
		goto L31
	} else {
		goto L37
	}
L34:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+4)))
	if v89 != v40 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v91 == int32(112) {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	goto L32
L38:
	;
	goto L9
L39:
	;
	goto L7
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L16
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v128 <= v129 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L42
L44:
	;
	F_createarc(m, l0, int32(112), int32(-2), l3, l4)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L16
	} else {
		goto L64
	}
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v131 == int32(0) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v153 == int32(0) {
		goto L44
	} else {
		goto L56
	}
L48:
	;
	v139 = v131
	goto L49
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	if v144 != l4 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L44
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	if v152 != 0 {
		v139 = v152
		goto L49
	} else {
		goto L55
	}
L52:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)))
	if v146 != int32(_a_F_rainbow_1) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v149 == int32(112) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	goto L50
L56:
	;
	v161 = v153
	goto L57
L57:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	if v166 != l3 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L44
L59:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	if v174 != 0 {
		v161 = v174
		goto L57
	} else {
		goto L63
	}
L60:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+4)))
	if v168 != int32(_a_F_rainbow_1) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v171 == int32(112) {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	goto L58
L64:
	;
	goto L1
}
func F_recordMultipleDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v235 int32
	_ = v235
	v5 = int32(0)
	if l2 <= v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_recordMultipleDependencies[0]))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v26 = int32(2340)
	if base.Ui32(v26) <= base.Ui32(l2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = v26
	goto L8
L7:
	;
	v29 = l2
	goto L8
L8:
	;
	v32 = F_palloc(m, v29<<(uint(int32(2))%32))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v35 = l1
	v40 = v5
	v42 = v5
	v44 = v5
	goto L11
L10:
	;
	if v184 != 0 {
		goto L42
	} else {
		goto L43
	}
L11:
	;
	v50 = v35
	v54 = int32(0)
	v57 = v42
	v59 = v44
	goto L13
L12:
	;
	if v163 <= int32(0) {
		v184 = v40
		v185 = v164
		goto L10
	} else {
		goto L36
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	goto L16
L14:
	;
	goto L12
L15:
	;
	v170 = v59 + int32(1)
	if v170 != l2 {
		v50 = v50 + int32(12)
		v54 = v163
		v57 = v164
		v59 = v170
		goto L13
	} else {
		goto L35
	}
L16:
	;
	if base.B2i32(base.B2i32(v63 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_recordMultipleDependencies_0)) < base.Ui32(v64)) == int32(0))&((base.B2i32(v63 != int32(2615))|base.B2i32(v64 != int32(2200)))&base.B2i32(v63 != int32(1262))) != 0 {
		v163 = v54
		v164 = v57
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if v29 <= v57 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	m.T0[v100].(func(*base.Module, int32))(m, v97)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L23
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v32+v54<<(uint(int32(2))%32))))
	v97 = v86
	v98 = v57
	goto L18
L20:
	;
	goto L21
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v92 = F_MakeTupleTableSlot(m, v90, int32(_a_F_recordMultipleDependencies_1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+v54<<(uint(int32(2))%32)))) = v92
	v97 = v92
	v98 = v57 + int32(1)
	goto L18
L23:
	;
	v105 = v32 + v54<<(uint(int32(2))%32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+16)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+24)) = base.I32_extend8_s(l3)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v135 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+20))
	base.MemoryFill(m, v136, int32(0), v135)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)))
	v142 = v140 & int32(_a_F_recordMultipleDependencies_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)) = uint16(v142)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*uint16)(unsafe.Add(mBase, uint32(v139)+6)) = uint16(v145)
	goto L27
L27:
	;
	v148 = v54 + int32(1)
	if v148 != v29 {
		v163 = v148
		v164 = v98
		goto L15
	} else {
		goto L28
	}
L28:
	;
	if v40 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v152 = F_CatalogOpenIndexes(m, v24)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v154 = v40
	goto L31
L31:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, v24, v32, v29, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	v154 = v152
	goto L31
L33:
	;
	v160 = v59 + int32(1)
	if v160 != l2 {
		v35 = v50 + int32(12)
		v40 = v154
		v42 = v98
		v44 = v160
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v184 = v154
	v185 = v98
	goto L10
L35:
	;
	goto L14
L36:
	;
	if v40 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v176 = F_CatalogOpenIndexes(m, v24)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v178 = v40
	goto L39
L39:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, v24, v32, v163, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L41
	}
L40:
	;
	v178 = v176
	goto L39
L41:
	;
	v184 = v178
	v185 = v164
	goto L10
L42:
	;
	F_CatalogCloseIndexes(m, v184)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_relation_close(m, v24, int32(3))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	if int32(0) < v185 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v198 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	F_pfree(m, v32)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L54
	}
L50:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v32+v198<<(uint(int32(2))%32))))
	F_ExecDropSingleTupleTableSlot(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	goto L49
L52:
	;
	v218 = v198 + int32(1)
	if v218 != v185 {
		v198 = v218
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L1
}
func F_record_image_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v586 int32
	_ = v586
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	v25 = m.G0
	v27 = v25 - int32(80)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v35 = F_pg_detoast_datum(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v39 = F_lookup_rowtype_tupdesc(m, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v44 = F_lookup_rowtype_tupdesc(m, v42, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = v30
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v49
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+68)) = uint16(v49)
	v53 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v53
	v55 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = int32(base.Ui32(v47) >> (uint(v55) % 32))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v49
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+48)) = uint16(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = int32(base.Ui32(v58) >> (uint(v55) % 32))
	if v46 < v41 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v70 = v41
	goto L8
L7:
	;
	v70 = v46
	goto L8
L8:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	if v72 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v37 != v95 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v83 = F_MemoryContextAlloc(m, v78, v70<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v75 < v70 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v94 = v72
	v95 = v77
	goto L9
L13:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v83
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v89 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v88)+12)) = v89
	v94 = v88
	v95 = int32(0)
	goto L9
L14:
	;
	v151 = F_palloc(m, v41<<(uint(int32(2))%32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L30
	}
L15:
	;
	v104 = v94 + int32(20)
	v108 = v70 << (uint(int32(2)) % 32)
	if v104&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v108)) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v97 != v38 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v99 != v42 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	if v101 == v43 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v37
	goto L14
L21:
	;
	if v108 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v108 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L24:
	;
	v118 = v94 + v108 + int32(20)
	v120 = v94 + int32(24)
	if base.Ui32(v120) < base.Ui32(v118) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v122 = v118
	goto L27
L26:
	;
	v122 = v120
	goto L27
L27:
	;
	v129 = (v122-v94-int32(21))&int32(-4) + int32(4)
	if v129 == int32(0) {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	base.MemoryFill(m, v104, int32(0), v129)
	goto L20
L29:
	;
	base.MemoryFill(m, v104, int32(0), v108)
	goto L20
L30:
	;
	v153 = F_palloc(m, v41)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_heap_deform_tuple(m, v27+int32(60), v39, v151, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v161 = F_palloc(m, v46<<(uint(int32(2))%32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v163 = F_palloc(m, v46)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_heap_deform_tuple(m, v27+int32(40), v44, v161, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v167 = int32(0)
	v171 = base.B2i32(v167 < v41)
	if v167 < v41 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L168
	}
L37:
	;
	F_pfree(m, v151)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L148
	}
L38:
	;
	if base.B2i32(v555 != v41)|base.B2i32(v554 != v46) != 0 {
		goto L36
	} else {
		goto L147
	}
L39:
	;
	v181 = v178
	v182 = v167
	v183 = base.B2i32(v167 < v46)
	v185 = v171
	v187 = v179
	goto L44
L40:
	;
	v172 = int32(0)
	v178 = v172
	v179 = v172
	goto L39
L41:
	;
	goto L42
L42:
	;
	v174 = int32(0)
	if v46 <= v174 {
		v554 = v174
		v555 = v167
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v178 = v174
	v179 = v174
	goto L39
L44:
	;
	if v185&int32(1) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v554 = v538
	v555 = v539
	goto L38
L46:
	;
	v550 = base.B2i32(v538 < v46)
	v551 = base.B2i32(v539 < v41)
	if v550|v551 != 0 {
		v181 = v538
		v182 = v539
		v183 = v550
		v185 = v551
		v187 = v543
		goto L44
	} else {
		goto L146
	}
L47:
	;
	if v183&int32(1) == int32(0) {
		v554 = v181
		v555 = v182
		goto L38
	} else {
		goto L50
	}
L48:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v208<<(uint(int32(4))%32)+v182*int32(100))+111)))
	if v215 != int32(1) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v538 = v181
	v539 = v182 + int32(1)
	v543 = v187
	goto L46
L50:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v230 = v44 + v224<<(uint(int32(4))%32) + v181*int32(100)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+111)))
	if v231 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v538 = v181 + int32(1)
	v539 = v182
	v543 = v187
	goto L46
L52:
	;
	goto L53
L53:
	;
	if v185&int32(1) == int32(0) {
		v554 = v181
		v555 = v182
		goto L38
	} else {
		goto L54
	}
L54:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v246 = v39 + v240<<(uint(int32(4))%32) + v182*int32(100)
	v248 = v246 + int32(20)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
	v251 = v230 + int32(88)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v249 == v252 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v532 = int32(1)
	v538 = v181 + v532
	v539 = v182 + v532
	v543 = v187 + v532
	goto L46
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L143
	}
L57:
	;
	v254 = int32(1)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v163))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v153))))
	if v258 == v254 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L137
	}
L60:
	;
	if v256&int32(1) == int32(0) {
		v586 = v254
		goto L37
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v265 = int32(-1)
	if v256&int32(1) != 0 {
		v586 = v265
		goto L37
	} else {
		goto L64
	}
L63:
	;
	goto L55
L64:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+82)))
	if v268 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v271 = int32(2)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v151+v182<<(uint(v271)%32))))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v161+v181<<(uint(v271)%32))))
	if v274 == v278 {
		goto L55
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+72)))
	if int32(0) < v284 {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	if base.Ui32(v274) < base.Ui32(v278) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v283 = int32(-1)
	goto L71
L70:
	;
	v283 = int32(1)
	goto L71
L71:
	;
	v586 = v283
	goto L37
L72:
	;
	if v468 < int32(0) {
		v586 = v265
		goto L37
	} else {
		goto L135
	}
L73:
	;
	v287 = int32(2)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v151+v182<<(uint(v287)%32))))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v161+v181<<(uint(v287)%32))))
	if base.Ui32(int32(4)) <= base.Ui32(v284) {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	goto L75
L75:
	;
	if v284 != int32(-1) {
		goto L56
	} else {
		goto L94
	}
L76:
	;
	v468 = v356
	goto L72
L77:
	;
	v356 = int32(0)
	goto L76
L78:
	;
	v330 = v325
	v331 = v326
	v332 = v327
	goto L88
L79:
	;
	if (v290|v294)&int32(3) != 0 {
		v325 = v290
		v326 = v294
		v327 = v284
		goto L78
	} else {
		goto L82
	}
L80:
	;
	v318 = v290
	v319 = v294
	v320 = v284
	goto L81
L81:
	;
	if v320 == int32(0) {
		goto L77
	} else {
		goto L87
	}
L82:
	;
	v302 = v290
	v303 = v294
	v304 = v284
	goto L83
L83:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	if v307 != v308 {
		v325 = v302
		v326 = v303
		v327 = v304
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v318 = v313
	v319 = v311
	v320 = v315
	goto L81
L85:
	;
	v310 = int32(4)
	v311 = v303 + v310
	v313 = v302 + v310
	v315 = v304 - v310
	if base.Ui32(int32(3)) < base.Ui32(v315) {
		v302 = v313
		v303 = v311
		v304 = v315
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v325 = v318
	v326 = v319
	v327 = v320
	goto L78
L88:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v335 == v336 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v356 = v335 - v336
	goto L76
L90:
	;
	v338 = int32(1)
	v343 = v332 - v338
	if v343 != 0 {
		v330 = v330 + v338
		v331 = v331 + v338
		v332 = v343
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	goto L77
L94:
	;
	v361 = v151 + v182<<(uint(int32(2))%32)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v363 = F_toast_raw_datum_size(m, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v367 = v161 + v181<<(uint(int32(2))%32)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v369 = F_toast_raw_datum_size(m, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v372 = base.B2i32(base.Ui32(v363) < base.Ui32(v369))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v375 = F_pg_detoast_datum_packed(m, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v378 = F_pg_detoast_datum_packed(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v380 = int32(1)
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	if v382&v380 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v385 = v380
	goto L101
L100:
	;
	v385 = int32(4)
	goto L101
L101:
	;
	v386 = v375 + v385
	v387 = int32(1)
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v389&v387 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v392 = v387
	goto L104
L103:
	;
	v392 = int32(4)
	goto L104
L104:
	;
	v393 = v378 + v392
	if base.Ui32(v363) < base.Ui32(v369) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v394 = v363
	goto L107
L106:
	;
	v394 = v369
	goto L107
L107:
	;
	v395 = int32(4)
	v396 = v394 - v395
	if base.Ui32(v395) <= base.Ui32(v396) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	if v459 != v375 {
		goto L126
	} else {
		goto L127
	}
L109:
	;
	v458 = int32(0)
	goto L108
L110:
	;
	v432 = v427
	v433 = v428
	v434 = v429
	goto L120
L111:
	;
	if (v386|v393)&int32(3) != 0 {
		v427 = v386
		v428 = v393
		v429 = v396
		goto L110
	} else {
		goto L114
	}
L112:
	;
	v420 = v386
	v421 = v393
	v422 = v396
	goto L113
L113:
	;
	if v422 == int32(0) {
		goto L109
	} else {
		goto L119
	}
L114:
	;
	v404 = v386
	v405 = v393
	v406 = v396
	goto L115
L115:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	if v409 != v410 {
		v427 = v404
		v428 = v405
		v429 = v406
		goto L110
	} else {
		goto L117
	}
L116:
	;
	v420 = v415
	v421 = v413
	v422 = v417
	goto L113
L117:
	;
	v412 = int32(4)
	v413 = v405 + v412
	v415 = v404 + v412
	v417 = v406 - v412
	if base.Ui32(int32(3)) < base.Ui32(v417) {
		v404 = v415
		v405 = v413
		v406 = v417
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v427 = v420
	v428 = v421
	v429 = v422
	goto L110
L120:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	if v437 == v438 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v458 = v437 - v438
	goto L108
L122:
	;
	v440 = int32(1)
	v445 = v434 - v440
	if v445 != 0 {
		v432 = v432 + v440
		v433 = v433 + v440
		v434 = v445
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	goto L109
L126:
	;
	F_pfree(m, v375)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if v458 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v463 = v458
	goto L132
L131:
	;
	v463 = base.B2i32(base.Ui32(v369) < base.Ui32(v363)) - v372
	goto L132
L132:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	if v378 == v464 {
		v468 = v463
		goto L72
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v378)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v468 = v463
	goto L72
L135:
	;
	if v468 == int32(0) {
		goto L55
	} else {
		goto L136
	}
L136:
	;
	v586 = int32(1)
	goto L37
L137:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v248)+68))
	v489 = F_format_type_be(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v492 = F_format_type_be(m, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v187 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v492
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v489
	F_errmsg(m, int32(_a_F_record_image_cmp_0), v27+int32(16))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_record_image_cmp_1), int32(1474), int32(_a_F_record_image_cmp_2))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	v513 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v513
	F_errmsg_internal(m, int32(_a_F_record_image_cmp_3), v27)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_record_image_cmp_1), int32(1538), int32(_a_F_record_image_cmp_2))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	goto L45
L147:
	;
	v586 = int32(0)
	goto L37
L148:
	;
	F_pfree(m, v153)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_pfree(m, v161)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_pfree(m, v163)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if int32(0) <= v613 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_DecrTupleDescRefCount(m, v39)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if int32(0) <= v618 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L154
L156:
	;
	F_DecrTupleDescRefCount(m, v44)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v623 != v30 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L158
L160:
	;
	F_pfree(m, v30)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v627 != v35 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L162
L164:
	;
	F_pfree(m, v35)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	m.G0 = v27 + int32(80)
	return v586
L167:
	;
	goto L166
L168:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(_a_F_record_image_cmp_4), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_record_image_cmp_1), int32(1568), int32(_a_F_record_image_cmp_2))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_record_lt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2) >> (uint(int32(31)) % 32))
	}
}
func F_record_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v242 int32
	_ = v242
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_stack_depth(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = int32(0)
	if base.B2i32(v22 == int32(2249))&base.B2i32(v23 < v31) == v31 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v499 = F_heap_form_tuple(m, v36, v109, v111)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L96
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L92
	}
L5:
	;
	if v114 == int32(0) {
		goto L3
	} else {
		goto L91
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L87
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L83
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L77
	}
L9:
	;
	v36 = F_lookup_rowtype_tupdesc(m, v22, v23)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L73
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v40 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v109 = F_palloc(m, v38<<(uint(int32(2))%32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	if v61 == v22 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v51 = F_MemoryContextAlloc(m, v46, v38*int32(44)+int32(12))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v43 != v38 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v60 = v40
	v61 = v45
	goto L14
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v51
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
	v60 = v56
	v61 = int32(0)
	goto L14
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v63 == v23 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v68 = v38 * int32(44)
	v70 = v68 + int32(12)
	if v60&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v70)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v22
	goto L13
L24:
	;
	if v70 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v70 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L27:
	;
	v82 = v60 + v68 + int32(12)
	v84 = v60 + int32(4)
	if base.Ui32(v84) < base.Ui32(v82) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v86 = v82
	goto L30
L29:
	;
	v86 = v84
	goto L30
L30:
	;
	v91 = (v60^int32(-1)+v86)&int32(-4) + int32(4)
	if v91 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	base.MemoryFill(m, v60, int32(0), v91)
	goto L23
L32:
	;
	base.MemoryFill(m, v60, int32(0), v70)
	goto L23
L33:
	;
	v111 = F_palloc(m, v38)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v114 = F_pq_getmsgint(m, v24, int32(4))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v38 <= int32(0) {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v119 = v38 & int32(3)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v121 = int32(4)
	v123 = v36 + v120<<(uint(v121)%32)
	v124 = int32(0)
	if base.Ui32(v121) <= base.Ui32(v38) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v222 != v114 {
		v449 = v222
		goto L4
	} else {
		goto L48
	}
L38:
	;
	v132 = v124
	v133 = v124
	v146 = v2
	goto L41
L39:
	;
	v175 = v124
	v176 = v124
	goto L40
L40:
	;
	v192 = v175
	v193 = v176
	v204 = v2
	goto L45
L41:
	;
	v149 = v123 + v132*int32(100)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+111)))
	v151 = int32(1)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+211)))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+311)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+411)))
	v165 = v133 + (v150 ^ v151) + (v154 ^ v151) + (v158 ^ v151) + (v162 ^ v151)
	v166 = int32(4)
	v167 = v132 + v166
	v169 = v146 + v166
	if v169 != v38&int32(2147483644) {
		v132 = v167
		v133 = v165
		v146 = v169
		goto L41
	} else {
		goto L43
	}
L42:
	;
	if v119 == int32(0) {
		v222 = v165
		goto L37
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v175 = v167
	v176 = v165
	goto L40
L45:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v192*int32(100))+111)))
	v211 = int32(1)
	v213 = v193 + (v210 ^ v211)
	v217 = v204 + v211
	if v217 != v119 {
		v192 = v192 + v211
		v193 = v213
		v204 = v217
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v222 = v213
	goto L37
L47:
	;
	goto L46
L48:
	;
	v242 = int32(0)
	goto L49
L49:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v263 = v36 + v257<<(uint(int32(4))%32) + v242*int32(100)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+111)))
	if v264 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L3
L51:
	;
	v358 = v242 + int32(1)
	if v358 != v38 {
		v242 = v358
		goto L49
	} else {
		goto L72
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109+v242<<(uint(int32(2))%32)))) = int32(0)
	v273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v242+v111))) = uint8(v273)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v263)+88))
	v277 = F_pq_getmsgint(m, v24, int32(4))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v280 = int32(_a_F_record_recv_0)
	if base.B2i32(base.B2i32(v275 == v277)|base.B2i32(base.Ui32(v280) < base.Ui32(v275)) == int32(0))&base.B2i32(base.Ui32(v277) <= base.Ui32(v280)) != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v289 = F_pq_getmsgint(m, v24, int32(4))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v289 < int32(-1) {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v293-v294 < v289 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v299 = v60 + int32(12) + v242*int32(44)
	if v289 == int32(-1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v242+v111))) = uint8(v315)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if v275 != v320 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v315 = int32(1)
	v317 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v289 + v294
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v294 + v306
	v315 = int32(0)
	v317 = v20 - int32(-64)
	goto L60
L64:
	;
	F_getTypeBinaryInputInfo(m, v275, v299+int32(4), v299+int32(8))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v299)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v263+int32(20))+76))
	v345 = F_ReceiveFunctionCall(m, v299+int32(16), v317, v341, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+20))
	F_fmgr_info_cxt(m, v328, v299+int32(16), v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299))) = v275
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109+v242<<(uint(int32(2))%32)))) = v345
	if v317 == int32(0) {
		goto L51
	} else {
		goto L70
	}
L70:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	if v350 != v289 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L51
L72:
	;
	goto L50
L73:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_record_recv_1), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(509), int32(_a_F_record_recv_3))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v385 = F_format_type_extended(m, v277, int32(-1), int32(2))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v389 = F_format_type_extended(m, v275, int32(-1), int32(2))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v242 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v277
	F_errmsg(m, int32(_a_F_record_recv_4), v20)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(606), int32(_a_F_record_recv_3))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_record_recv_5), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(613), int32(_a_F_record_recv_3))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v242 + int32(1)
	F_errmsg(m, int32(_a_F_record_recv_6), v20+int32(32))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(661), int32(_a_F_record_recv_3))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v449 = int32(0)
	goto L4
L92:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v114
	F_errmsg(m, int32(_a_F_record_recv_7), v20+int32(48))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(559), int32(_a_F_record_recv_3))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v502 = F_palloc(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	if v504 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v499)+16))
	base.MemoryCopy(m, v502, v505, v504)
	goto L100
L99:
	;
	goto L100
L100:
	;
	F_pfree(m, v499)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_pfree(m, v109)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_pfree(m, v111)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if int32(0) <= v513 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	F_DecrTupleDescRefCount(m, v36)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v518 = F_HeapTupleHeaderGetDatum(m, v502)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	m.G0 = v20 + int32(80)
	return v518
}
func F_recurse_pushdown_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = l0
	goto L2
L1:
	;
	m.G0 = v8 + int32(16)
	return v58
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v15 != int32(142) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L14
	}
L4:
	;
	goto L3
L5:
	;
	if v15 != int32(63) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v35 == int32(3) {
		v58 = v34
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)-int32(4))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v30 = F_subquery_is_pushdown_safe(m, v29, l1, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v58 = v30
	goto L1
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v39 = F_recurse_pushdown_safe(m, v38, l1, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if v39 == int32(0) {
		v58 = v34
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v10 = v43
	goto L2
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v48
	F_errmsg_internal(m, int32(_a_F_recurse_pushdown_safe_0), v8)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_recurse_pushdown_safe_1), int32(3712), int32(_a_F_recurse_pushdown_safe_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_recvfrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v7 = m.Env.X__syscall_recvfrom(m, l0, l1, l2, l3, l4, l5)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _c_F_recvfrom[0])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_redirect_elem_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	F_appendStringInfo(m, l0, int32(_a_F_redirect_elem_desc_0), v7)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_regclassout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(0) {
		v14 = F_pstrdup(m, int32(_a_F_regclassout_0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v53 = v14
			m.G0 = v8 + int32(16)
			return v53
		}
	} else {
		v19 = F_SearchSysCache1(m, int32(57), v10)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v25 = v23 + int32(4)
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_regclassout[0]))
				if v27 == int32(0) {
					v30 = F_pstrdup(m, v25)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v19)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v53 = v30
							m.G0 = v8 + int32(16)
							return v53
						}
					}
				} else {
					v34 = F_RelationIsVisible(m, v10)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							v40 = int32(0)
							v41 = F_quote_qualified_identifier(m, v40, v25)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v53 = v41
									m.G0 = v8 + int32(16)
									return v53
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
							v38 = F_get_namespace_name(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = v38
								v41 = F_quote_qualified_identifier(m, v40, v25)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v19)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v53 = v41
										m.G0 = v8 + int32(16)
										return v53
									}
								}
							}
						}
					}
				}
			} else {
				v46 = F_palloc(m, int32(64))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					v51 = F_pg_snprintf(m, v46, int32(64), int32(_a_F_regclassout_1), v8)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = v46
						m.G0 = v8 + int32(16)
						return v53
					}
				}
			}
		}
	}
}
func F_regex_selectivity_sub(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 float64
	_ = v59
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 float64
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v138 float64
	_ = v138
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v145 int32
	_ = v145
	var v149 float64
	_ = v149
	var v152 float64
	_ = v152
	var v155 float64
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 float64
	_ = v167
	var v174 float64
	_ = v174
	var v177 float64
	_ = v177
	v5 = int32(0)
	F_check_stack_depth(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(0)
L2:
	;
	v14 = float64(1)
	if l1 <= int32(0) {
		v167 = v14
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v174 = float64(1)
	if base.F64_gt(v167, v174) != 0 {
		goto L68
	} else {
		goto L69
	}
L4:
	;
	v19 = v14
	v21 = v5
	v22 = v5
	v24 = v5
	goto L5
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v21))))
	if v27 == int32(40) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v167 = v155
	goto L3
L7:
	;
	v163 = v157 + int32(1)
	if v163 < l1 {
		v19 = v155
		v21 = v163
		v22 = v158
		v24 = v160
		goto L5
	} else {
		goto L67
	}
L8:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v35 = int32(0)
	if base.B2i32(v27 != int32(41))|base.B2i32(v22 <= v35) == v35 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v30 = v24
	goto L13
L12:
	;
	v30 = v21
	goto L13
L13:
	;
	v155 = v19
	v157 = v21
	v158 = v22 + int32(1)
	v160 = v30
	goto L7
L14:
	;
	v41 = v22 - int32(1)
	if v41 != 0 {
		v155 = v19
		v157 = v21
		v158 = v41
		v160 = v24
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.B2i32(v27 != int32(124))|v22 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v43 = v24 + int32(1)
	v46 = F_regex_selectivity_sub(m, l0+v43, v21-v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v155 = base.F64_mul(v19, v46)
	v157 = v21
	v158 = int32(0)
	v160 = v24
	goto L7
L19:
	;
	v56 = v21 + int32(1)
	v59 = F_regex_selectivity_sub(m, l0+v56, l1-v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	switch v27 - int32(42) {
	case 0, 1, 21:
		goto L25
	case 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48:
		goto L23
	case 4:
		goto L26
	case 49:
		goto L27
	case 50:
		goto L24
	default:
		goto L28
	}
L22:
	;
	v167 = base.F64_add(v19, v59)
	goto L3
L23:
	;
	if v22 != 0 {
		goto L64
	} else {
		goto L65
	}
L24:
	;
	v145 = v21 + int32(1)
	if l1 <= v145 {
		v167 = v19
		goto L3
	} else {
		goto L60
	}
L25:
	;
	if v22 != 0 {
		goto L57
	} else {
		goto L58
	}
L26:
	;
	if v22 != 0 {
		goto L54
	} else {
		goto L55
	}
L27:
	;
	v97 = v21 + int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v97))))
	v101 = base.B2i32(v99 == int32(94))
	if v99 == int32(94) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	if v27 != int32(123) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	if l1 <= v21 {
		v87 = v21
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v22 != 0 {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v71 = v21
	goto L32
L32:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v71))))
	if v77 == int32(125) {
		v87 = v71
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v87 = l1
	goto L30
L34:
	;
	v81 = v71 + int32(1)
	if v81 != l1 {
		v71 = v81
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v93 = v19
	goto L38
L37:
	;
	v93 = base.F64_add(v19, v19)
	goto L38
L38:
	;
	v155 = v93
	v157 = v87
	v158 = v22
	v160 = v24
	goto L7
L39:
	;
	v102 = float64(0.75)
	goto L41
L40:
	;
	v102 = float64(0.25)
	goto L41
L41:
	;
	if v99 == int32(94) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v22 != 0 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	v105 = v21 + int32(2)
	goto L45
L44:
	;
	v105 = v97
	goto L45
L45:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v105))))
	v110 = v105 + base.B2i32(v107 == int32(93))
	if l1 <= v110 {
		v132 = v110
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v116 = v110
	goto L47
L47:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v116))))
	if v122 == int32(93) {
		v132 = v116
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v132 = l1
	goto L42
L49:
	;
	v126 = v116 + int32(1)
	if v126 < l1 {
		v116 = v126
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v138 = v19
	goto L53
L52:
	;
	v138 = base.F64_mul(v19, v102)
	goto L53
L53:
	;
	v155 = v138
	v157 = v132
	v158 = v22
	v160 = v24
	goto L7
L54:
	;
	v141 = v19
	goto L56
L55:
	;
	v141 = base.F64_mul(v19, float64(0.9))
	goto L56
L56:
	;
	v155 = v141
	v157 = v21
	v158 = v22
	v160 = v24
	goto L7
L57:
	;
	v143 = v19
	goto L59
L58:
	;
	v143 = base.F64_add(v19, v19)
	goto L59
L59:
	;
	v155 = v143
	v157 = v21
	v158 = v22
	v160 = v24
	goto L7
L60:
	;
	if v22 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v149 = v19
	goto L63
L62:
	;
	v149 = base.F64_mul(v19, float64(0.2))
	goto L63
L63:
	;
	v155 = v149
	v157 = v145
	v158 = v22
	v160 = v24
	goto L7
L64:
	;
	v152 = v19
	goto L66
L65:
	;
	v152 = base.F64_mul(v19, float64(0.2))
	goto L66
L66:
	;
	v155 = v152
	v157 = v21
	v158 = v22
	v160 = v24
	goto L7
L67:
	;
	goto L6
L68:
	;
	v177 = v174
	goto L70
L69:
	;
	v177 = v167
	goto L70
L70:
	;
	return v177
}
func F_regexnesel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13991(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_register_dirty_segment(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v181 int64
	_ = v181
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v206 int64
	_ = v206
	var v211 int32
	_ = v211
	v2 = l1
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v14
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+10)) = uint16(v2)
	v17 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v17
	v23 = F_RegisterSyncRequest(m, v8+int32(8), v4, v4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		if v23 == int32(0) {
			v29 = F_errstart(m, int32(14), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				if v29 != 0 {
					F_errmsg_internal(m, int32(_a_F_register_dirty_segment_0), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_register_dirty_segment_1), int32(1522), int32(_a_F_register_dirty_segment_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[0])))
							v44 = m.G0
							v46 = v44 - int32(16)
							m.G0 = v46
							if v41 != 0 {
								F___clock_gettime(m, int32(1), v46)
								mBase = m.M
								v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
								v55 = v50 + v51*int64(1000000000)
							} else {
								v55 = int64(0)
							}
							m.G0 = v46 + int32(16)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v61 = F_FileSync(m, v59, int32(167772182))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								if int32(0) <= v61 {
									v95 = int32(1)
									v97 = int64(0)
									v101 = m.G0
									v103 = v101 - int32(16)
									m.G0 = v103
									if v55 != v97 {
										F___clock_gettime(m, int32(1), v103)
										mBase = m.M
										v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
										v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
										v114 = v109 + (v110*int64(1000000000) - v55)
										v158 = int32(0)
										v164 = int32(200)
										v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
										*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
										v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
										if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
											v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
											*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
											v185 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
										} else {
										}
									} else {
									}
									v201 = int32(200)
									v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
									v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
									F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
									mBase = m.M
									v211 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
									m.G0 = v103 + int32(16)
									m.G0 = v8 + int32(32)
									return
								} else {
									v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[9])))
									if v68 != 0 {
										v69 = int32(21)
									} else {
										v69 = int32(23)
									}
									v71 = F_errstart(m, v69, int32(0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										if v71 == int32(0) {
											v95 = int32(1)
											v97 = int64(0)
											v101 = m.G0
											v103 = v101 - int32(16)
											m.G0 = v103
											if v55 != v97 {
												F___clock_gettime(m, int32(1), v103)
												mBase = m.M
												v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
												v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
												v114 = v109 + (v110*int64(1000000000) - v55)
												v158 = int32(0)
												v164 = int32(200)
												v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
												*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
												v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
												if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
													v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
													*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
													v185 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
													*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
												} else {
												}
											} else {
											}
											v201 = int32(200)
											v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
											*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
											v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
											*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
											F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
											mBase = m.M
											v211 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
											m.G0 = v103 + int32(16)
											m.G0 = v8 + int32(32)
											return
										} else {
											F_errcode_for_file_access(m)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v79 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[10]))
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77*int32(48))+32))
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v83
												F_errmsg(m, int32(_a_F_register_dirty_segment_4), v8)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_register_dirty_segment_1), int32(1530), int32(_a_F_register_dirty_segment_2))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														v95 = int32(1)
														v97 = int64(0)
														v101 = m.G0
														v103 = v101 - int32(16)
														m.G0 = v103
														if v55 != v97 {
															F___clock_gettime(m, int32(1), v103)
															mBase = m.M
															v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
															v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
															v114 = v109 + (v110*int64(1000000000) - v55)
															v158 = int32(0)
															v164 = int32(200)
															v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
															*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
															v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
															if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
																v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
																*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
																v185 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
																*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
															} else {
															}
														} else {
														}
														v201 = int32(200)
														v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
														*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
														v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
														*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
														F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
														mBase = m.M
														v211 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
														*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
														m.G0 = v103 + int32(16)
														m.G0 = v8 + int32(32)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[0])))
					v44 = m.G0
					v46 = v44 - int32(16)
					m.G0 = v46
					if v41 != 0 {
						F___clock_gettime(m, int32(1), v46)
						mBase = m.M
						v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
						v55 = v50 + v51*int64(1000000000)
					} else {
						v55 = int64(0)
					}
					m.G0 = v46 + int32(16)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v61 = F_FileSync(m, v59, int32(167772182))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						if int32(0) <= v61 {
							v95 = int32(1)
							v97 = int64(0)
							v101 = m.G0
							v103 = v101 - int32(16)
							m.G0 = v103
							if v55 != v97 {
								F___clock_gettime(m, int32(1), v103)
								mBase = m.M
								v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
								v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
								v114 = v109 + (v110*int64(1000000000) - v55)
								v158 = int32(0)
								v164 = int32(200)
								v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
								*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
								v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
								if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
									v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
									v185 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
								} else {
								}
							} else {
							}
							v201 = int32(200)
							v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
							*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
							v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
							*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
							F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
							mBase = m.M
							v211 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
							*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
							m.G0 = v103 + int32(16)
							m.G0 = v8 + int32(32)
							return
						} else {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[9])))
							if v68 != 0 {
								v69 = int32(21)
							} else {
								v69 = int32(23)
							}
							v71 = F_errstart(m, v69, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								if v71 == int32(0) {
									v95 = int32(1)
									v97 = int64(0)
									v101 = m.G0
									v103 = v101 - int32(16)
									m.G0 = v103
									if v55 != v97 {
										F___clock_gettime(m, int32(1), v103)
										mBase = m.M
										v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
										v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
										v114 = v109 + (v110*int64(1000000000) - v55)
										v158 = int32(0)
										v164 = int32(200)
										v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
										*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
										v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
										if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
											v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
											*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
											v185 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
										} else {
										}
									} else {
									}
									v201 = int32(200)
									v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
									v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
									F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
									mBase = m.M
									v211 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
									m.G0 = v103 + int32(16)
									m.G0 = v8 + int32(32)
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v79 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[10]))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77*int32(48))+32))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v83
										F_errmsg(m, int32(_a_F_register_dirty_segment_4), v8)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_register_dirty_segment_1), int32(1530), int32(_a_F_register_dirty_segment_2))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v95 = int32(1)
												v97 = int64(0)
												v101 = m.G0
												v103 = v101 - int32(16)
												m.G0 = v103
												if v55 != v97 {
													F___clock_gettime(m, int32(1), v103)
													mBase = m.M
													v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
													v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
													v114 = v109 + (v110*int64(1000000000) - v55)
													v158 = int32(0)
													v164 = int32(200)
													v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
													*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
													v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
													if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
														v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
														*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
														v185 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
														*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
													} else {
													}
												} else {
												}
												v201 = int32(200)
												v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
												*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
												v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
												*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
												F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
												mBase = m.M
												v211 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
												*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
												m.G0 = v103 + int32(16)
												m.G0 = v8 + int32(32)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_regoperout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == int32(0) {
		v15 = F_pstrdup(m, int32(_a_F_regoperout_0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v86 = v15
			m.G0 = v9 + int32(32)
			return v86
		}
	} else {
		v20 = F_SearchSysCache1(m, int32(40), v11)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
				v24 = v22 + v23
				v26 = v24 + int32(4)
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_regoperout[0]))
				if v28 == int32(0) {
					v31 = F_pstrdup(m, v26)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v20)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v86 = v31
							m.G0 = v9 + int32(32)
							return v86
						}
					}
				} else {
					v35 = F_makeString(m, v26)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v35
						v42 = F_list_make1_impl(m, int32(1), v9+int32(24))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = int32(0)
							v46 = F_OpernameGetCandidates(m, v42, v44, v44)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								if v46 == int32(0) {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
									v58 = F_get_namespace_name(m, v57)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v60 = F_quote_identifier(m, v58)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = F_strlen(m, v60)
											mBase = m.M
											v63 = F_strlen(m, v26)
											mBase = m.M
											v67 = F_palloc(m, v62+v63+int32(2))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v26
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v60
												v74 = F_pg_sprintf(m, v67, int32(_a_F_regoperout_1), v9+int32(16))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v20)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														v86 = v67
														m.G0 = v9 + int32(32)
														return v86
													}
												}
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
									if v50 != 0 {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
										v58 = F_get_namespace_name(m, v57)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = F_quote_identifier(m, v58)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												v62 = F_strlen(m, v60)
												mBase = m.M
												v63 = F_strlen(m, v26)
												mBase = m.M
												v67 = F_palloc(m, v62+v63+int32(2))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v26
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v60
													v74 = F_pg_sprintf(m, v67, int32(_a_F_regoperout_1), v9+int32(16))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v20)
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return int32(0)
														} else {
															v86 = v67
															m.G0 = v9 + int32(32)
															return v86
														}
													}
												}
											}
										}
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
										if v51 != v11 {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
											v58 = F_get_namespace_name(m, v57)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												v60 = F_quote_identifier(m, v58)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int32(0)
												} else {
													v62 = F_strlen(m, v60)
													mBase = m.M
													v63 = F_strlen(m, v26)
													mBase = m.M
													v67 = F_palloc(m, v62+v63+int32(2))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v26
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v60
														v74 = F_pg_sprintf(m, v67, int32(_a_F_regoperout_1), v9+int32(16))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int32(0)
														} else {
															F_ReleaseCatCache(m, v20)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return int32(0)
															} else {
																v86 = v67
																m.G0 = v9 + int32(32)
																return v86
															}
														}
													}
												}
											}
										} else {
											v53 = F_pstrdup(m, v26)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return int32(0)
												} else {
													v86 = v53
													m.G0 = v9 + int32(32)
													return v86
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v79 = F_palloc(m, int32(64))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					v84 = F_pg_snprintf(m, v79, int32(64), int32(_a_F_regoperout_2), v9)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						v86 = v79
						m.G0 = v9 + int32(32)
						return v86
					}
				}
			}
		}
	}
}
func F_regprocout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == int32(0) {
		v15 = F_pstrdup(m, int32(_a_F_regprocout_0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v74 = v15
			m.G0 = v9 + int32(16)
			return v74
		}
	} else {
		v20 = F_SearchSysCache1(m, int32(47), v11)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
				v24 = v22 + v23
				v26 = v24 + int32(4)
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_regprocout[0]))
				if v28 == int32(0) {
					v31 = F_pstrdup(m, v26)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v20)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v74 = v31
							m.G0 = v9 + int32(16)
							return v74
						}
					}
				} else {
					v35 = F_makeString(m, v26)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35
						v42 = F_list_make1_impl(m, int32(1), v9+int32(8))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v45 = int32(0)
							v50 = F_FuncnameGetCandidates(m, v42, int32(-1), v45, v45, v45, v45, v45)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
									v59 = F_get_namespace_name(m, v58)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v61 = v59
										v62 = F_quote_qualified_identifier(m, v61, v26)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v20)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v74 = v62
												m.G0 = v9 + int32(16)
												return v74
											}
										}
									}
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
									if v54 != 0 {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
										v59 = F_get_namespace_name(m, v58)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v61 = v59
											v62 = F_quote_qualified_identifier(m, v61, v26)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v74 = v62
													m.G0 = v9 + int32(16)
													return v74
												}
											}
										}
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
										if v56 == v11 {
											v61 = int32(0)
											v62 = F_quote_qualified_identifier(m, v61, v26)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v74 = v62
													m.G0 = v9 + int32(16)
													return v74
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
											v59 = F_get_namespace_name(m, v58)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v61 = v59
												v62 = F_quote_qualified_identifier(m, v61, v26)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v20)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														v74 = v62
														m.G0 = v9 + int32(16)
														return v74
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v67 = F_palloc(m, int32(64))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					v72 = F_pg_snprintf(m, v67, int32(64), int32(_a_F_regprocout_1), v9)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = v67
						m.G0 = v9 + int32(16)
						return v74
					}
				}
			}
		}
	}
}
func F_regprocrecv(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_int4recv(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_regtypeout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v9 == int32(0) {
		v13 = F_pstrdup(m, int32(_a_F_regtypeout_0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v45 = v13
			m.G0 = v7 + int32(16)
			return v45
		}
	} else {
		v18 = F_SearchSysCache1(m, int32(82), v9)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_regtypeout[0]))
				if v21 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
					v29 = F_pstrdup(m, v24+v25+int32(4))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v18)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v45 = v29
							m.G0 = v7 + int32(16)
							return v45
						}
					}
				} else {
					v33 = F_format_type_be(m, v9)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v18)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v45 = v33
							m.G0 = v7 + int32(16)
							return v45
						}
					}
				}
			} else {
				v38 = F_palloc(m, int32(64))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
					v43 = F_pg_snprintf(m, v38, int32(64), int32(_a_F_regtypeout_1), v7)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v38
						m.G0 = v7 + int32(16)
						return v45
					}
				}
			}
		}
	}
}
func F_relmap_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	if base.Ui32(l0) < base.Ui32(int32(16)) {
		v6 = int32(_a_F_relmap_identify_0)
	} else {
		v6 = int32(0)
	}
	return v6
}
func F_relmap_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v5 = m.G0
	v7 = v5 - int32(544)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	v12 = v10 & int32(240)
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		if v16 != int32(524) {
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v73
				F_errmsg_internal(m, int32(_a_F_relmap_redo_0), v7)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_relmap_redo_1), int32(1111), int32(_a_F_relmap_redo_2))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = v7 + int32(20)
			base.MemoryCopy(m, v20, v15+int32(12), int32(524))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v27 = F_GetDatabasePath(m, v25, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_relmap_redo[0]))
				v34 = F_LWLockAcquire(m, v30+int32(3200), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int32(0)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					F_write_relmap_file(m, v20, v36, int32(1), v36, v39, v40, v27)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_relmap_redo[0]))
						F_LWLockRelease(m, v44+int32(3200))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_pfree(m, v27)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								m.G0 = v7 + int32(544)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v12
			F_errmsg_internal(m, int32(_a_F_relmap_redo_3), v7+int32(16))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_relmap_redo_1), int32(1140), int32(_a_F_relmap_redo_2))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_remove_dbtablespaces(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v17 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+188))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	m.T0[v156].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L44
	}
L2:
	;
	return
L3:
	;
	v19 = int32(0)
	v21 = F_table_beginscan_catalog(m, v17, v19, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = F_heap_getnext(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v23 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = v23
	v30 = int32(0)
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)))
	if v40 != int32(1664) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v80 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L9:
	;
	v43 = F_GetDatabasePath(m, l0, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L13
	}
L10:
	;
	v80 = v30
	goto L11
L11:
	;
	v81 = F_heap_getnext(m, v21)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L26
	}
L12:
	;
	F_pfree(m, v43)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L25
	}
L13:
	;
	v49 = F___fstatat(m, int32(-100), v43, v13+int32(16), int32(256))
	mBase = m.M
	goto L14
L14:
	;
	if v49 < int32(0) {
		v76 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v52&int32(_a_F_remove_dbtablespaces_0) != int32(_a_F_remove_dbtablespaces_1) {
		v76 = v30
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v57 = F_rmtree(m, v43)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	v74 = F_lappend_oid(m, v30, v40)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L24
	}
L18:
	;
	if v57 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v61 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	if v61 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v43
	F_errmsg(m, int32(_a_F_remove_dbtablespaces_2), v13)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_remove_dbtablespaces_3), int32(3040), int32(_a_F_remove_dbtablespaces_4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	v76 = v74
	goto L12
L25:
	;
	v80 = v76
	goto L11
L26:
	;
	if v81 != 0 {
		v28 = v81
		v30 = v80
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v87 = v85 << (uint(int32(2)) % 32)
	v88 = F_palloc(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	goto L1
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if int32(0) < v90 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	F_XLogBeginInsert(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L38
	}
L35:
	;
	v105 = v95 << (uint(int32(2)) % 32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107+v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v88+v105))) = v109
	v112 = v95 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v112 < v113 {
		v95 = v112
		goto L35
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	goto L36
L38:
	;
	F_XLogRegisterData(m, v13+int32(16), int32(8))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_XLogRegisterData(m, v88, v87)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v138 = F_XLogInsert(m, int32(4), int32(33))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_list_free(m, v80)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_pfree(m, v88)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L1
L44:
	;
	F_relation_close(m, v17, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	m.G0 = v13 + int32(112)
	return
}
func F_remove_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13846(m, l0, l1, l2, int32(1053))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_rename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v3 = int32(-100)
	v5 = m.Env.X__syscall_renameat(m, v3, l0, v3, l1)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v5) {
		*(*int32)(unsafe.Add(mBase, _c_F_rename[0])) = int32(0) - v5
		v13 = int32(-1)
	} else {
		v13 = v5
	}
	return v13
}
func F_rename_constraint_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v122 int32
	_ = v122
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L7
	} else {
		goto L60
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L56
	}
L3:
	;
	v36 = F_SearchSysCache1(m, int32(19), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L12
	}
L4:
	;
	v22 = F_get_domain_constraint_oid(m, l2, l3, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v25 = F_relation_open(m, l1, int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	v34 = int32(0)
	v35 = v22
	goto L3
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	F_renameatt_check(m, l1, v27, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v32 = F_get_relation_constraint_oid(m, l1, l3, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v34 = v25
	v35 = v32
	goto L3
L12:
	;
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	v40 = v38 + v39
	if l1 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L53
	}
L16:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v40)+88))
	if v139 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+72)))
	switch v43 - int32(99) {
	case 0, 11:
		goto L18
	default:
		goto L16
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+106)))
	if v46 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if l5 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+104)))
	if l6 < v122 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v50 = F_find_all_inheritors(m, l1, int32(8), v16+int32(-4))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if l6 != 0 {
		goto L20
	} else {
		goto L37
	}
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v56 = int32(0)
	goto L25
L25:
	;
	v69 = int32(0)
	if v50 == v69 {
		v79 = v69
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v52 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v73 <= v56 {
		v79 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v79 = v75 + v56<<(uint(int32(2))%32)
	goto L27
L30:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if base.B2i32(v79 == int32(0))|base.B2i32(v84 <= v56) != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v87 == int32(0) {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if l1 != v90 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v94 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v87+v56<<(uint(int32(2))%32))))
	F_rename_constraint_internal(m, v16+int32(-16), v90, v94, l3, l4, v94, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v56 = v56 + int32(1)
	goto L25
L36:
	;
	goto L35
L37:
	;
	v105 = F_find_inheritance_children(m, l1, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	if v105 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L20
L40:
	;
	goto L16
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	F_ReleaseCatCache(m, v36)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L47
	}
L42:
	;
	F_RenameConstraintById(m, v35, l4)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L7
	} else {
		goto L46
	}
L43:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+72)))
	v144 = v142 - int32(112)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v144))|base.B2i32(int32(1)<<(uint(v144)%32)&int32(289) == int32(0)) != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_RenameRelationInternal(m, v139, l4, int32(0), int32(1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	goto L41
L47:
	;
	if v34 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_CacheInvalidateRelcache(m, v34)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	m.G0 = v18 - int32(-64)
	return
L51:
	;
	F_relation_close(m, v34, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v35
	F_errmsg_internal(m, int32(_a_F_rename_constraint_internal_0), v18)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_rename_constraint_internal_1), int32(4083), int32(_a_F_rename_constraint_internal_2))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l3
	F_errmsg(m, int32(_a_F_rename_constraint_internal_3), v16+int32(-32))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_rename_constraint_internal_1), int32(_a_F_rename_constraint_internal_4), int32(_a_F_rename_constraint_internal_2))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l3
	F_errmsg(m, int32(_a_F_rename_constraint_internal_5), v16+int32(-48))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_rename_constraint_internal_1), int32(_a_F_rename_constraint_internal_6), int32(_a_F_rename_constraint_internal_2))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_reorderqueue_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v7 = F_pairingheap_remove_first(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v12 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = int32(0)
	v18 = v12
	goto L6
L4:
	;
	goto L5
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	F_pfree(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v17))))
	if v22 != 0 {
		v34 = v18
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v36 = v17 + int32(1)
	if v36 < v34 {
		v17 = v36
		v18 = v34
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v17))))
	if v25 != 0 {
		v34 = v18
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v17<<(uint(int32(2))%32))))
	F_pfree(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v34 = v33
	goto L8
L12:
	;
	goto L7
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	F_pfree(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_pfree(m, v7)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	return v11
}
func F_repalloc0(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if base.Ui32(l1) <= base.Ui32(l2) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13&int32(15)*int32(36))+uint32(_c_F_repalloc0[0])))
		v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = l2 - l1
			if v23 != 0 {
				base.MemoryFill(m, v19+l1, int32(0), v23)
			} else {
			}
			m.G0 = v7 + int32(16)
			return v19
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
			F_errmsg_internal(m, int32(_a_F_repalloc0_0), v7)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_repalloc0_1), int32(1655), int32(_a_F_repalloc0_2))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_repalloc_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6&int32(15)*int32(36))+uint32(_c_F_repalloc_extended[0])))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(2))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_repalloc_huge(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6&int32(15)*int32(36))+uint32(_c_F_repalloc_huge[0])))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_reparameterize_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 float64
	_ = v157
	var v158 int32
	_ = v158
	var v159 float64
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 float64
	_ = v304
	var v305 float64
	_ = v305
	var v306 float64
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 float64
	_ = v315
	var v319 float64
	_ = v319
	var v327 float64
	_ = v327
	var v333 float64
	_ = v333
	var v339 float64
	_ = v339
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 float64
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v389 float64
	_ = v389
	var v398 float64
	_ = v398
	var v402 float64
	_ = v402
	var v406 int32
	_ = v406
	var v408 float64
	_ = v408
	var v410 float64
	_ = v410
	var v413 float64
	_ = v413
	var v416 float64
	_ = v416
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	v5 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = v13
	goto L3
L2:
	;
	v14 = v5
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = int32(0)
	if v14 == v16 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return v429
L5:
	;
	return v422
L6:
	;
	if v69 == int32(0) {
		v422 = v5
		goto L5
	} else {
		goto L20
	}
L7:
	;
	v69 = int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	if l2 == int32(0) {
		v62 = v16
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v69 = v62
	goto L6
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v26 < v25 {
		v62 = v16
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(1)
	if v25 <= v28 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v31 = v28
	goto L15
L14:
	;
	v31 = v25
	goto L15
L15:
	;
	v32 = int32(8)
	v37 = int32(0)
	goto L16
L16:
	;
	v44 = v37 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14+v32+v44)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2+v32+v44)))
	v51 = v46 & (v48 ^ int32(-1))
	v53 = base.B2i32(v51 == int32(0))
	if v51 != 0 {
		v62 = v53
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v62 = v53
	goto L10
L18:
	;
	v55 = v37 + int32(1)
	if v55 != v31 {
		v37 = v55
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v72 - int32(331) {
	case 0:
		goto L24
	default:
		v422 = v5
		goto L5
	case 3:
		goto L23
	case 8:
		goto L29
	case 9:
		goto L28
	case 10, 11:
		goto L27
	case 13:
		goto L26
	case 16:
		goto L25
	case 29:
		goto L22
	case 30:
		goto L21
	}
L21:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v350 = F_reparameterize_path(m, l0, v349, l2, l3)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L30
	} else {
		goto L80
	}
L22:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v273 = F_reparameterize_path(m, l0, v272, l2, l3)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L30
	} else {
		goto L70
	}
L23:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v213 == int32(0) {
		v262 = v5
		v263 = v5
		goto L53
	} else {
		goto L54
	}
L24:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v189 != int32(279) {
		v422 = v5
		goto L5
	} else {
		goto L49
	}
L25:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v157 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v158)+56))
	v161 = F_palloc0(m, int32(80))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L30
	} else {
		goto L43
	}
L26:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v135 = F_palloc0(m, int32(80))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L30
	} else {
		goto L40
	}
L27:
	;
	v120 = F_palloc0(m, int32(112))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L30
	} else {
		goto L37
	}
L28:
	;
	v99 = F_palloc0(m, int32(72))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L30
	} else {
		goto L34
	}
L29:
	;
	v76 = F_palloc0(m, int32(72))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v76))) = int64(1455993913623)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v83
	v85 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+20)) = uint8(v87)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v85
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+64)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = v87
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+21)) = uint8(v90)
	F_cost_seqscan(m, v76, l0, v15, v85)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v429 = v76
	goto L4
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v99))) = int64(1460288880919)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v104
	v106 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+20)) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+16)) = v106
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+64)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = v108
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+21)) = uint8(v111)
	F_cost_samplescan(m, v99, l0, v15, v106)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v429 = v99
	goto L4
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(280)
	base.MemoryCopy(m, v120, l1, int32(112))
	v126 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v126
	F_cost_index(m, v120, l0, l3, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L30
	} else {
		goto L39
	}
L39:
	;
	return v120
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = int64(1477468750106)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v140
	v142 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L30
	} else {
		goto L41
	}
L41:
	;
	v144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+20)) = uint8(v144)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+16)) = v142
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+72)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v135)+64)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v135)+24)) = v144
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+21)) = uint8(v147)
	F_cost_bitmap_heap_scan(m, v135, l0, v15, v142, v133, l3)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v429 = v135
	goto L4
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v161))) = int64(1490353651999)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+12)) = v166
	v168 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v170 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+20)) = uint8(v170)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = v168
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	if v173 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+21)))
	v178 = v176
	goto L47
L46:
	;
	v178 = int32(0)
	goto L47
L47:
	;
	v180 = v178 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+21)) = uint8(v180)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+72)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v161)+64)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = v182
	F_cost_subqueryscan(m, v161, l0, v15, v168, base.F64_eq(v157, v159))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v429 = v161
	goto L4
L49:
	;
	v193 = F_palloc0(m, int32(72))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L30
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v193))) = int64(1421634175255)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = v198
	v200 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L30
	} else {
		goto L51
	}
L51:
	;
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+20)) = uint8(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v193)+16)) = v200
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+64)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v193)+24)) = v202
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+21)) = uint8(v205)
	F_cost_resultscan(m, v193, l0, v15, v200)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L30
	} else {
		goto L52
	}
L52:
	;
	v429 = v193
	goto L4
L53:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	v269 = F_create_append_path(m, l0, v15, v262, v263, v265, l2, v266, v267, float64(-1))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L30
	} else {
		goto L69
	}
L54:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v216 <= int32(0) {
		v262 = v5
		v263 = v5
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v226 = int32(0)
	v228 = v5
	v229 = v5
	goto L56
L56:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v226<<(uint(int32(2))%32))))
	v236 = F_reparameterize_path(m, l0, v235, l2, l3)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L30
	} else {
		goto L58
	}
L57:
	;
	v262 = v248
	v263 = v249
	goto L53
L58:
	;
	if v236 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	return int32(0)
L60:
	;
	goto L61
L61:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v226 < v242 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v251 = v226 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v251 < v252 {
		v226 = v251
		v228 = v248
		v229 = v249
		goto L56
	} else {
		goto L68
	}
L63:
	;
	v244 = F_lappend(m, v228, v236)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L30
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v246 = F_lappend(m, v229, v236)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L30
	} else {
		goto L67
	}
L66:
	;
	v248 = v244
	v249 = v229
	goto L62
L67:
	;
	v248 = v228
	v249 = v246
	goto L62
L68:
	;
	goto L57
L69:
	;
	return v269
L70:
	;
	if v273 == int32(0) {
		v422 = v5
		goto L5
	} else {
		goto L71
	}
L71:
	;
	v278 = F_palloc0(m, int32(80))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L30
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = int64(1546188226853)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+12)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v278)+20)) = uint8(v286)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = v285
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	if v289 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+21)))
	v294 = v292
	goto L75
L74:
	;
	v294 = int32(0)
	goto L75
L75:
	;
	v296 = v294 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v278)+21)) = uint8(v296)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v273)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v298
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v273)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+72)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v278)+64)) = v300
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v273)+40))
	v304 = *(*float64)(unsafe.Add(mBase, uint32(v273)+48))
	v305 = *(*float64)(unsafe.Add(mBase, uint32(v273)+56))
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v273)+32))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+32))
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_reparameterize_path[0]))
	*(*float64)(unsafe.Add(mBase, uint32(v278)+32)) = v306
	v315 = *(*float64)(unsafe.Add(mBase, _c_F_reparameterize_path[1]))
	v319 = base.F64_add(base.F64_mul(base.F64_add(v315, v315), v306), base.F64_sub(v305, v304))
	v327 = base.F64_mul(v306, base.F64_convert_i32_u((v308+int32(7))&int32(-8)+int32(24)))
	if base.F64_gt(v327, base.F64_convert_i32_u(v312<<(uint(int32(10))%32))) != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v429 = v278
	goto L4
L77:
	;
	v333 = *(*float64)(unsafe.Add(mBase, _c_F_reparameterize_path[2]))
	v339 = base.F64_add(base.F64_mul(v333, base.F64_ceil(base.F64_mul(v327, float64(0.0001220703125)))), v319)
	goto L79
L78:
	;
	v339 = v319
	goto L79
L79:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reparameterize_path[3])))
	*(*float64)(unsafe.Add(mBase, uint32(v278)+56)) = base.F64_add(v304, v339)
	*(*float64)(unsafe.Add(mBase, uint32(v278)+48)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v278)+40)) = v303 + (v341 ^ int32(1))
	goto L76
L80:
	;
	if v350 == int32(0) {
		v422 = v5
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v354 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+85)))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+84)))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v360 = F_palloc0(m, int32(104))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L30
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v360))) = int64(1550483194150)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+12)) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v350)+16))
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v360)+20)) = uint8(v368)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+16)) = v367
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	if v371 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+21)))
	v376 = v374
	goto L85
L84:
	;
	v376 = int32(0)
	goto L85
L85:
	;
	v378 = v376 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v360)+21)) = uint8(v378)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v350)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+24)) = v380
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v350)+64))
	*(*uint8)(unsafe.Add(mBase, uint32(v360)+85)) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, uint32(v360)+84)) = uint8(v356)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+80)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v360)+76)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v360)+72)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v360)+64)) = v382
	v389 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v354)&int64(9223372036854775807)))|base.F64_gt(v354, v389) != 0 {
		v402 = v389
		goto L87
	} else {
		goto L88
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+96)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(v360)+88)) = v402
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v350)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+40)) = v406
	v408 = *(*float64)(unsafe.Add(mBase, uint32(v350)+48))
	v410 = *(*float64)(unsafe.Add(mBase, _c_F_reparameterize_path[4]))
	*(*float64)(unsafe.Add(mBase, uint32(v360)+48)) = base.F64_add(v408, v410)
	v413 = *(*float64)(unsafe.Add(mBase, uint32(v350)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v360)+56)) = base.F64_add(v410, v413)
	v416 = *(*float64)(unsafe.Add(mBase, uint32(v350)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v360)+32)) = v416
	v422 = v360
	goto L5
L87:
	;
	goto L86
L88:
	;
	v398 = float64(1)
	if base.F64_le(v354, v398) != 0 {
		v402 = v398
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v402 = base.F64_nearest(v354)
	goto L87
}
func F_replace_vars_in_jointree(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(63) {
	case 0:
		goto L7
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L4
	}
L3:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v134 = F_replace_rte_variables(m, v130, v16, int32(0), int32(851), l1, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L36
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L14
	} else {
		goto L33
	}
L5:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_replace_vars_in_jointree(m, v96, l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L14
	} else {
		goto L27
	}
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v60 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v15 == v16 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+v15<<(uint(int32(2))%32)-int32(4))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+124)))
	if v28 != int32(1) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	switch v31 {
	case 0:
		goto L3
	case 1:
		goto L13
	default:
		goto L1
	case 3:
		goto L12
	case 4:
		goto L11
	case 5:
		goto L10
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v57 = F_replace_rte_variables(m, v53, v16, int32(0), int32(851), l1, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L18
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v50 = F_replace_rte_variables(m, v46, v16, int32(0), int32(851), l1, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L17
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v43 = F_replace_rte_variables(m, v39, v16, int32(0), int32(851), l1, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L16
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	v36 = F_replace_rte_variables(m, v32, v16, int32(1), int32(851), l1, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v36
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v43
	goto L1
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = v50
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v57
	goto L1
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v92 = F_replace_rte_variables(m, v87, v88, int32(0), int32(851), l1, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L26
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v63 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v70 = v3
	goto L22
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v70<<(uint(int32(2))%32))))
	F_replace_vars_in_jointree(m, v75, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L14
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	v79 = v70 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v79 < v80 {
		v70 = v79
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92
	goto L1
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_replace_vars_in_jointree(m, v99, l1)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v102 == int32(2) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(2)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v112 = F_replace_rte_variables(m, v107, v108, int32(0), int32(851), l1, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v95
	goto L1
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v120
	F_errmsg_internal(m, int32(_a_F_replace_vars_in_jointree_0), v8)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_replace_vars_in_jointree_1), int32(2613), int32(_a_F_replace_vars_in_jointree_2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L14
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
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v134
	goto L1
}
func F_report_corruption(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+100)))
	v17 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+84)))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+124)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v21 = F_Int64GetDatum(m, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = base.I32_extend16_s(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v21
		v28 = int32(base.Ui32(v18) >> (uint(int32(15)) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v28)
		v30 = F_cstring_to_text(m, l1)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v30
			F_pfree(m, l1)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v39 = F_heap_form_tuple(m, v15, v12+int32(16), v12+int32(12))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_tuplestore_puttuple(m, v14, v39)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v43)
						m.G0 = v12 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_report_invalid_encoding_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	if l2 < l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = l2
	goto L3
L2:
	;
	v13 = l3
	goto L3
L3:
	;
	if int32(0) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = int32(8)
	if v16 <= v13 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L19
	}
L7:
	;
	v19 = v16
	goto L9
L8:
	;
	v19 = v13
	goto L9
L9:
	;
	v27 = int32(0)
	v28 = v10 + int32(32)
	goto L10
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v27))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v33
	v38 = F_pg_sprintf(m, v28, int32(_a_F_report_invalid_encoding_int_0), v10+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L6
L12:
	;
	return
L13:
	;
	v40 = v38 + v28
	if v27 < v19-int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = F_pg_sprintf(m, v40, int32(_a_F_report_invalid_encoding_int_1), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	v47 = v40
	goto L16
L16:
	;
	v49 = v27 + int32(1)
	if v49 != v19 {
		v27 = v49
		v28 = v47
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v47 = v44 + v40
	goto L16
L18:
	;
	goto L11
L19:
	;
	F_errcode(m, int32(17301634))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_report_invalid_encoding_int[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v10 + int32(32)
	F_errmsg(m, int32(_a_F_report_invalid_encoding_int_2), v10)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_report_invalid_encoding_int_3), int32(1853), int32(_a_F_report_invalid_encoding_int_4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_report_triggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 float64
	_ = v103
	var v104 float64
	_ = v104
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 float64
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 float64
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v16 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(80)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 <= int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = v16
	v34 = v4
	goto L5
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v40 = v37 + v34*int32(416)
	F_InstrEndLoop(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	return
L8:
	;
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v40)+216))
	if base.F64_ne(v43, float64(0)) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v46 = int32(0)
	F_ExplainOpenGroup(m, int32(_a_F_report_triggers_0), v46, int32(1), l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v161 = v34 + int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v161 < v163 {
		v28 = v162
		v34 = v161
		goto L5
	} else {
		goto L55
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	v56 = v36 + v34*int32(60)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	if v57 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = F_get_constraint_name(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	v60 = v46
	goto L15
L15:
	;
	v62 = v53 + int32(4)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v63 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v60 = v58
	goto L15
L17:
	;
	if v60 != 0 {
		goto L50
	} else {
		goto L51
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v68 = int32(0)
	if v67|base.B2i32(v60 == v68) == v68 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	goto L20
L20:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	F_ExplainPropertyText(m, int32(_a_F_report_triggers_1), v120, l2)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L39
	}
L21:
	;
	if l1 != 0 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v60
	F_appendStringInfo(m, v85, int32(_a_F_report_triggers_2), v14+int32(48))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L29
	}
L23:
	;
	F_appendStringInfoString(m, v66, int32(_a_F_report_triggers_0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v76
	F_appendStringInfo(m, v66, int32(_a_F_report_triggers_3), v14-int32(-64))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L7
	} else {
		goto L27
	}
L26:
	;
	goto L22
L27:
	;
	if v60 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	goto L21
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v62
	F_appendStringInfo(m, v92, int32(_a_F_report_triggers_4), v14+int32(32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v100 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v40)+208))
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v40)+216))
	*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v104
	*(*float64)(unsafe.Add(mBase, uint32(v14))) = base.F64_mul(v103, float64(1000))
	F_appendStringInfo(m, v99, int32(_a_F_report_triggers_5), v14)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v112 = *(*float64)(unsafe.Add(mBase, uint32(v40)+216))
	*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = v112
	F_appendStringInfo(m, v99, int32(_a_F_report_triggers_6), v14+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L38
	}
L37:
	;
	goto L17
L38:
	;
	goto L17
L39:
	;
	if v60 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_ExplainPropertyText(m, int32(_a_F_report_triggers_7), v60, l2)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_ExplainPropertyText(m, int32(_a_F_report_triggers_8), v62, l2)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v129 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v40)+208))
	F_ExplainPropertyFloat(m, int32(_a_F_report_triggers_9), int32(_a_F_report_triggers_10), base.F64_mul(v134, float64(1000)), int32(3), l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v141 = int32(0)
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v40)+216))
	F_ExplainPropertyFloat(m, int32(_a_F_report_triggers_11), v141, v142, v141, l2)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	goto L17
L50:
	;
	F_pfree(m, v60)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_ExplainCloseGroup(m, int32(_a_F_report_triggers_0), int32(1), l2)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	goto L11
L55:
	;
	goto L6
}
func F_reserveAllocatedDesc(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[0]))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[1]))
	if v8 < v6 {
		v56 = int32(1)
		return v56
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[2]))
		if v11 == int32(0) {
			v16 = F_emscripten_builtin_malloc(m, int32(192))
			mBase = m.M
			if v16 != 0 {
				v46 = int32(16)
				v47 = v16
				*(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[0])) = v46
				*(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[2])) = v47
				v56 = int32(1)
				return v56
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_reserveAllocatedDesc_0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_reserveAllocatedDesc_1), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_reserveAllocatedDesc_2), int32(2597), int32(_a_F_reserveAllocatedDesc_3))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v35 = int32(0)
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[3]))
			v39 = base.I32_div_s(v37, int32(3))
			if v39 <= v6 {
				v56 = v35
			} else {
				v43 = F_emscripten_builtin_realloc(m, v11, v39*int32(12))
				mBase = m.M
				if v43 == int32(0) {
					v56 = v35
				} else {
					v46 = v39
					v47 = v43
					*(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[0])) = v46
					*(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[2])) = v47
					v56 = int32(1)
				}
			}
			return v56
		}
	}
}
func F_resolve_anymultirange_from_others(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = F_getBaseType(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = F_get_range_multirange(m, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				if v11 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v40 = F_format_type_be(m, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v40
								F_errmsg(m, int32(_a_F_resolve_anymultirange_from_others_0), v6)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_resolve_anymultirange_from_others_1), int32(726), int32(_a_F_resolve_anymultirange_from_others_2))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v11
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_resolve_anymultirange_from_others_3), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_resolve_anymultirange_from_others_1), int32(730), int32(_a_F_resolve_anymultirange_from_others_2))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_restore(m *base.Module, l0 int32, l1 float32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
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
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v776 int32
	_ = v776
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = base.F64_promote_f32(l1)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_restore[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+56)) = uint8(v18)
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_restore[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v21
	v24 = *(*int64)(unsafe.Add(mBase, _c_F_restore[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v24
	v27 = *(*int64)(unsafe.Add(mBase, _c_F_restore[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v27
	v29 = int32(6)
	if base.Ui32(v29) <= base.Ui32(l2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = v29
	goto L3
L2:
	;
	v33 = l2
	goto L3
L3:
	;
	if l2 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = v29
	goto L6
L5:
	;
	v36 = v33
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v36 - int32(1)
	v43 = F_pg_sprintf(m, l0, int32(_a_F_restore_0), v11+int32(-48))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v47 = int32(101)
	v48 = F___strchrnul(m, l0, v47)
	mBase = m.M
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v50 == v47 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v776 = F_strlen(m, l0)
	mBase = m.M
	m.G0 = v13 - int32(-64)
	return v776
L10:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L14
	}
L11:
	;
	v54 = v48
	goto L13
L12:
	;
	v54 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v62 = v54 + int32(1)
	goto L16
L15:
	;
	if v106 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	v67 = v62 + int32(1)
	v68 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	v69 = F___isspace(m, v68)
	mBase = m.M
	if v69 != 0 {
		v62 = v67
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v70 = int32(1)
	switch v68&int32(255) - int32(43) {
	case 0:
		v76 = v70
		goto L20
	default:
		v78 = v68
		v79 = v62
		v80 = v70
		goto L19
	case 2:
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v81 = int32(0)
	v83 = v78 - int32(48)
	if base.Ui32(v83) <= base.Ui32(int32(9)) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67))))
	v78 = v77
	v79 = v67
	v80 = v76
	goto L19
L21:
	;
	v76 = int32(0)
	goto L20
L22:
	;
	v86 = v81
	v87 = v83
	v88 = v79
	goto L25
L23:
	;
	v100 = v81
	goto L24
L24:
	;
	if v80 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v90 = int32(10)
	v92 = v86*v90 - v87
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v88)+1)))
	v97 = v93 - int32(48)
	if base.Ui32(v97) < base.Ui32(v90) {
		v86 = v92
		v87 = v97
		v88 = v88 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v100 = v92
	goto L24
L27:
	;
	goto L26
L28:
	;
	v106 = int32(0) - v100
	goto L30
L29:
	;
	v106 = v100
	goto L30
L30:
	;
	goto L15
L31:
	;
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v109)
	goto L9
L32:
	;
	goto L33
L33:
	;
	v112 = v106 >> (uint(int32(31)) % 32)
	if base.Ui32(int32(4)) < base.Ui32(v106^v112-v112) {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v117 = int32(10)
	v120 = l0 + base.F32_lt(l1, float32(0))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v121 != int32(101) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = v117
	v128 = v121
	v130 = v120
	v131 = int32(0)
	goto L38
L36:
	;
	v154 = v117
	goto L37
L37:
	;
	if int32(0) < v106 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-32)+v126))) = uint8(v128)
	v141 = base.B2i32(v128&int32(255) == int32(46))
	if v128&int32(255) == int32(46) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v142 != 0 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v142 = v126
	goto L42
L41:
	;
	v142 = v131
	goto L42
L42:
	;
	v144 = int32(1)
	v145 = v126 - v141 + v144
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v146 != int32(101) {
		v126 = v145
		v128 = v146
		v130 = v130 + v144
		v131 = v142
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v151 = v142
	goto L46
L45:
	;
	v151 = v145
	goto L46
L46:
	;
	v154 = v151
	goto L37
L47:
	;
	v164 = v154 + v106
	v166 = v164 - int32(10)
	if v36 <= v166 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v600 = v11 + int32(-32)
	v602 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v600+v36)+10)) = uint8(v602)
	v605 = v154 + v600 + v106
	v608 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v605-int32(1)))) = uint8(v608)
	if base.F32_lt(l1, float32(0)) != 0 {
		goto L159
	} else {
		goto L160
	}
L50:
	;
	v169 = v11 + int32(-32)
	v170 = v169 + v36
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+10)) = uint8(v171)
	if base.Ui32(int32(2)) <= base.Ui32(v36) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if int32(22) < v164 {
		goto L102
	} else {
		goto L103
	}
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v13)+51))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v175
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v13)+43))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+44)) = v177
	v179 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+43)) = uint8(v179)
	v181 = v170
	goto L55
L54:
	;
	v181 = v169
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v166 - int32(1)
	v188 = F_pg_sprintf(m, v181+int32(11), int32(_a_F_restore_1), v13)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	if base.F32_lt(l1, float32(0)) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v192 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+41)) = uint8(v192)
	v197 = v11 + int32(-32) | int32(9)
	if (v197^l0)&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	goto L59
L59:
	;
	v275 = v11 + int32(-32) | int32(10)
	if (v275^l0)&int32(3) != 0 {
		goto L84
	} else {
		goto L85
	}
L60:
	;
	goto L9
L61:
	;
	goto L60
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v252))) = uint8(v251)
	if v251&int32(255) == int32(0) {
		goto L61
	} else {
		goto L77
	}
L63:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v250 = v197
	v251 = v203
	v252 = l0
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v197&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v207 = v197
	v209 = l0
	goto L69
L67:
	;
	v221 = v197
	v223 = l0
	goto L68
L68:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v228 = int32(-2139062144)
	if (int32(16843008)-v225|v225)&v228 != v228 {
		v250 = v221
		v251 = v225
		v252 = v223
		goto L62
	} else {
		goto L73
	}
L69:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v209))) = uint8(v210)
	if v210 == int32(0) {
		goto L61
	} else {
		goto L71
	}
L70:
	;
	v221 = v217
	v223 = v215
	goto L68
L71:
	;
	v214 = int32(1)
	v215 = v209 + v214
	v217 = v207 + v214
	if v217&int32(3) != 0 {
		v207 = v217
		v209 = v215
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v233 = v221
	v234 = v225
	v235 = v223
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v234
	v237 = int32(4)
	v238 = v235 + v237
	v240 = v233 + v237
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	v245 = int32(-2139062144)
	if (int32(16843008)-v242|v242)&v245 == v245 {
		v233 = v240
		v234 = v242
		v235 = v238
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v250 = v240
	v251 = v242
	v252 = v238
	goto L62
L76:
	;
	goto L75
L77:
	;
	v259 = v250
	v261 = v252
	goto L78
L78:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)) = uint8(v262)
	v264 = int32(1)
	if v262 != 0 {
		v259 = v259 + v264
		v261 = v261 + v264
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L61
L80:
	;
	goto L79
L81:
	;
	goto L9
L82:
	;
	goto L81
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v330))) = uint8(v329)
	if v329&int32(255) == int32(0) {
		goto L82
	} else {
		goto L98
	}
L84:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v328 = v275
	v329 = v281
	v330 = l0
	goto L83
L85:
	;
	goto L86
L86:
	;
	if v275&int32(3) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v285 = v275
	v287 = l0
	goto L90
L88:
	;
	v299 = v275
	v301 = l0
	goto L89
L89:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	v306 = int32(-2139062144)
	if (int32(16843008)-v303|v303)&v306 != v306 {
		v328 = v299
		v329 = v303
		v330 = v301
		goto L83
	} else {
		goto L94
	}
L90:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	*(*uint8)(unsafe.Add(mBase, uint32(v287))) = uint8(v288)
	if v288 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L91:
	;
	v299 = v295
	v301 = v293
	goto L89
L92:
	;
	v292 = int32(1)
	v293 = v287 + v292
	v295 = v285 + v292
	if v295&int32(3) != 0 {
		v285 = v295
		v287 = v293
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v311 = v299
	v312 = v303
	v313 = v301
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v312
	v315 = int32(4)
	v316 = v313 + v315
	v318 = v311 + v315
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	v323 = int32(-2139062144)
	if (int32(16843008)-v320|v320)&v323 == v323 {
		v311 = v318
		v312 = v320
		v313 = v316
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v328 = v318
	v329 = v320
	v330 = v316
	goto L83
L97:
	;
	goto L96
L98:
	;
	v337 = v328
	v339 = v330
	goto L99
L99:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+1)) = uint8(v340)
	v342 = int32(1)
	if v340 != 0 {
		v337 = v337 + v342
		v339 = v339 + v342
		goto L99
	} else {
		goto L101
	}
L100:
	;
	goto L82
L101:
	;
	goto L100
L102:
	;
	v434 = v11 + int32(-32)
	v436 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v434+v36)+11)) = uint8(v436)
	v439 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v434+v164))) = uint8(v439)
	if base.F32_lt(l1, float32(0)) != 0 {
		goto L114
	} else {
		goto L115
	}
L103:
	;
	v352 = int32(23)
	v354 = v352 - v164
	v355 = int32(3)
	v356 = v354 & v355
	if base.Ui32(v355) <= base.Ui32(v164-int32(20)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v366 = v352
	v368 = int32(0)
	goto L107
L105:
	;
	v392 = v352
	goto L106
L106:
	;
	v403 = v392
	v405 = int32(0)
	goto L111
L107:
	;
	v376 = v11 + int32(-32) + v366
	v379 = int32(4)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v376-v379)))
	*(*int32)(unsafe.Add(mBase, uint32(v376-int32(3)))) = v381
	v384 = v366 - v379
	v386 = v368 + v379
	if v386 != v354&int32(-4) {
		v366 = v384
		v368 = v386
		goto L107
	} else {
		goto L109
	}
L108:
	;
	if v356 == int32(0) {
		goto L102
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	v392 = v384
	goto L106
L111:
	;
	v413 = v11 + int32(-32) + v403
	v414 = int32(1)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413-v414))))
	*(*uint8)(unsafe.Add(mBase, uint32(v413))) = uint8(v416)
	v421 = v405 + v414
	if v421 != v356 {
		v403 = v403 - v414
		v405 = v421
		goto L111
	} else {
		goto L113
	}
L112:
	;
	goto L102
L113:
	;
	goto L112
L114:
	;
	v443 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+41)) = uint8(v443)
	v446 = v434 | int32(9)
	if (v446^l0)&int32(3) != 0 {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	goto L116
L116:
	;
	v524 = v11 + int32(-32) | int32(10)
	if (v524^l0)&int32(3) != 0 {
		goto L141
	} else {
		goto L142
	}
L117:
	;
	goto L9
L118:
	;
	goto L117
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v501))) = uint8(v500)
	if v500&int32(255) == int32(0) {
		goto L118
	} else {
		goto L134
	}
L120:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	v499 = v446
	v500 = v452
	v501 = l0
	goto L119
L121:
	;
	goto L122
L122:
	;
	if v446&int32(3) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v456 = v446
	v458 = l0
	goto L126
L124:
	;
	v470 = v446
	v472 = l0
	goto L125
L125:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v477 = int32(-2139062144)
	if (int32(16843008)-v474|v474)&v477 != v477 {
		v499 = v470
		v500 = v474
		v501 = v472
		goto L119
	} else {
		goto L130
	}
L126:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v459)
	if v459 == int32(0) {
		goto L118
	} else {
		goto L128
	}
L127:
	;
	v470 = v466
	v472 = v464
	goto L125
L128:
	;
	v463 = int32(1)
	v464 = v458 + v463
	v466 = v456 + v463
	if v466&int32(3) != 0 {
		v456 = v466
		v458 = v464
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v482 = v470
	v483 = v474
	v484 = v472
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = v483
	v486 = int32(4)
	v487 = v484 + v486
	v489 = v482 + v486
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	v494 = int32(-2139062144)
	if (int32(16843008)-v491|v491)&v494 == v494 {
		v482 = v489
		v483 = v491
		v484 = v487
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v499 = v489
	v500 = v491
	v501 = v487
	goto L119
L133:
	;
	goto L132
L134:
	;
	v508 = v499
	v510 = v501
	goto L135
L135:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)) = uint8(v511)
	v513 = int32(1)
	if v511 != 0 {
		v508 = v508 + v513
		v510 = v510 + v513
		goto L135
	} else {
		goto L137
	}
L136:
	;
	goto L118
L137:
	;
	goto L136
L138:
	;
	goto L9
L139:
	;
	goto L138
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v579))) = uint8(v578)
	if v578&int32(255) == int32(0) {
		goto L139
	} else {
		goto L155
	}
L141:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	v577 = v524
	v578 = v530
	v579 = l0
	goto L140
L142:
	;
	goto L143
L143:
	;
	if v524&int32(3) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v534 = v524
	v536 = l0
	goto L147
L145:
	;
	v548 = v524
	v550 = l0
	goto L146
L146:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v555 = int32(-2139062144)
	if (int32(16843008)-v552|v552)&v555 != v555 {
		v577 = v548
		v578 = v552
		v579 = v550
		goto L140
	} else {
		goto L151
	}
L147:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v537)
	if v537 == int32(0) {
		goto L139
	} else {
		goto L149
	}
L148:
	;
	v548 = v544
	v550 = v542
	goto L146
L149:
	;
	v541 = int32(1)
	v542 = v536 + v541
	v544 = v534 + v541
	if v544&int32(3) != 0 {
		v534 = v544
		v536 = v542
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v560 = v548
	v561 = v552
	v562 = v550
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v561
	v564 = int32(4)
	v565 = v562 + v564
	v567 = v560 + v564
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v572 = int32(-2139062144)
	if (int32(16843008)-v569|v569)&v572 == v572 {
		v560 = v567
		v561 = v569
		v562 = v565
		goto L152
	} else {
		goto L154
	}
L153:
	;
	v577 = v567
	v578 = v569
	v579 = v565
	goto L140
L154:
	;
	goto L153
L155:
	;
	v586 = v577
	v588 = v579
	goto L156
L156:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v588)+1)) = uint8(v589)
	v591 = int32(1)
	if v589 != 0 {
		v586 = v586 + v591
		v588 = v588 + v591
		goto L156
	} else {
		goto L158
	}
L157:
	;
	goto L139
L158:
	;
	goto L157
L159:
	;
	v612 = int32(3)
	v613 = v605 - v612
	v614 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v613))) = uint8(v614)
	if (v613^l0)&v612 != 0 {
		goto L165
	} else {
		goto L166
	}
L160:
	;
	goto L161
L161:
	;
	v691 = v605 - int32(2)
	if (v691^l0)&int32(3) != 0 {
		goto L186
	} else {
		goto L187
	}
L162:
	;
	goto L9
L163:
	;
	goto L162
L164:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v670))) = uint8(v669)
	if v669&int32(255) == int32(0) {
		goto L163
	} else {
		goto L179
	}
L165:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	v668 = v613
	v669 = v621
	v670 = l0
	goto L164
L166:
	;
	goto L167
L167:
	;
	if v613&int32(3) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v625 = v613
	v627 = l0
	goto L171
L169:
	;
	v639 = v613
	v641 = l0
	goto L170
L170:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v639)))
	v646 = int32(-2139062144)
	if (int32(16843008)-v643|v643)&v646 != v646 {
		v668 = v639
		v669 = v643
		v670 = v641
		goto L164
	} else {
		goto L175
	}
L171:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	*(*uint8)(unsafe.Add(mBase, uint32(v627))) = uint8(v628)
	if v628 == int32(0) {
		goto L163
	} else {
		goto L173
	}
L172:
	;
	v639 = v635
	v641 = v633
	goto L170
L173:
	;
	v632 = int32(1)
	v633 = v627 + v632
	v635 = v625 + v632
	if v635&int32(3) != 0 {
		v625 = v635
		v627 = v633
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v651 = v639
	v652 = v643
	v653 = v641
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653))) = v652
	v655 = int32(4)
	v656 = v653 + v655
	v658 = v651 + v655
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v651)+4))
	v663 = int32(-2139062144)
	if (int32(16843008)-v660|v660)&v663 == v663 {
		v651 = v658
		v652 = v660
		v653 = v656
		goto L176
	} else {
		goto L178
	}
L177:
	;
	v668 = v658
	v669 = v660
	v670 = v656
	goto L164
L178:
	;
	goto L177
L179:
	;
	v677 = v668
	v679 = v670
	goto L180
L180:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v679)+1)) = uint8(v680)
	v682 = int32(1)
	if v680 != 0 {
		v677 = v677 + v682
		v679 = v679 + v682
		goto L180
	} else {
		goto L182
	}
L181:
	;
	goto L163
L182:
	;
	goto L181
L183:
	;
	goto L9
L184:
	;
	goto L183
L185:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v746))) = uint8(v745)
	if v745&int32(255) == int32(0) {
		goto L184
	} else {
		goto L200
	}
L186:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691))))
	v744 = v691
	v745 = v697
	v746 = l0
	goto L185
L187:
	;
	goto L188
L188:
	;
	if v691&int32(3) != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v701 = v691
	v703 = l0
	goto L192
L190:
	;
	v715 = v691
	v717 = l0
	goto L191
L191:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v722 = int32(-2139062144)
	if (int32(16843008)-v719|v719)&v722 != v722 {
		v744 = v715
		v745 = v719
		v746 = v717
		goto L185
	} else {
		goto L196
	}
L192:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701))))
	*(*uint8)(unsafe.Add(mBase, uint32(v703))) = uint8(v704)
	if v704 == int32(0) {
		goto L184
	} else {
		goto L194
	}
L193:
	;
	v715 = v711
	v717 = v709
	goto L191
L194:
	;
	v708 = int32(1)
	v709 = v703 + v708
	v711 = v701 + v708
	if v711&int32(3) != 0 {
		v701 = v711
		v703 = v709
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	v727 = v715
	v728 = v719
	v729 = v717
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v729))) = v728
	v731 = int32(4)
	v732 = v729 + v731
	v734 = v727 + v731
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	v739 = int32(-2139062144)
	if (int32(16843008)-v736|v736)&v739 == v739 {
		v727 = v734
		v728 = v736
		v729 = v732
		goto L197
	} else {
		goto L199
	}
L198:
	;
	v744 = v734
	v745 = v736
	v746 = v732
	goto L185
L199:
	;
	goto L198
L200:
	;
	v753 = v744
	v755 = v746
	goto L201
L201:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v755)+1)) = uint8(v756)
	v758 = int32(1)
	if v756 != 0 {
		v753 = v753 + v758
		v755 = v755 + v758
		goto L201
	} else {
		goto L203
	}
L202:
	;
	goto L184
L203:
	;
	goto L202
}
func F_restrict_and_check_grant(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v192 int64
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int64
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	v14 = m.G0
	v16 = v14 - int32(192)
	m.G0 = v16
	switch l6 - int32(6) {
	case 0:
		v58 = int64(167503724583)
		goto L1
	default:
		goto L3
	case 3:
		goto L11
	case 8:
		goto L5
	case 10, 11, 15, 43:
		goto L9
	case 13:
		goto L10
	case 16:
		goto L8
	case 21:
		goto L4
	case 30:
		goto L7
	case 31:
		goto L12
	case 35:
		goto L2
	case 36:
		goto L6
	}
L1:
	;
	if l1 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v58 = int64(70914205040767)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L17
	}
L4:
	;
	v58 = int64(52776558145536)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v58 = int64(2199023256064)
	goto L1
L7:
	;
	v58 = int64(3298534884096)
	goto L1
L8:
	;
	v58 = int64(25769803782)
	goto L1
L9:
	;
	v58 = int64(1099511628032)
	goto L1
L10:
	;
	v58 = int64(549755814016)
	goto L1
L11:
	;
	v58 = int64(15393162792448)
	goto L1
L12:
	;
	v58 = int64(1125281431814)
	goto L1
L13:
	;
	return int64(0)
L14:
	;
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_14), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(285), int32(_a_F_restrict_and_check_grant_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l6
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_12), v16)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(295), int32(_a_F_restrict_and_check_grant_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L13
	} else {
		goto L142
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L13
	} else {
		goto L138
	}
L22:
	;
	v209 = l3 & int64(base.Ui64(l1)>>(uint(int64(32))%64))
	if l0 != 0 {
		goto L86
	} else {
		goto L87
	}
L23:
	;
	switch l6 - int32(6) {
	case 0:
		goto L25
	default:
		goto L26
	case 3:
		goto L38
	case 8:
		goto L28
	case 10:
		goto L30
	case 11:
		goto L29
	case 13:
		goto L37
	case 15:
		goto L36
	case 16:
		goto L35
	case 21:
		goto L34
	case 30:
		goto L33
	case 31, 35:
		goto L39
	case 33:
		goto L32
	case 36:
		goto L31
	case 43:
		goto L27
	}
L24:
	;
	if v192 != int64(0) {
		goto L22
	} else {
		goto L78
	}
L25:
	;
	v184 = F_pg_class_aclmask_ext(m, l4, l5, v58, int32(1), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L13
	} else {
		goto L76
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L73
	}
L27:
	;
	v165 = F_object_aclmask_ext(m, int32(1247), l4, l5, v58, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L72
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L69
	}
L29:
	;
	v148 = F_object_aclmask_ext(m, int32(1417), l4, l5, v58, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L68
	}
L30:
	;
	v144 = F_object_aclmask_ext(m, int32(2328), l4, l5, v58, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L67
	}
L31:
	;
	v140 = F_object_aclmask_ext(m, int32(1213), l4, l5, v58, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L66
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L63
	}
L33:
	;
	v123 = F_object_aclmask_ext(m, int32(2615), l4, l5, v58, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L13
	} else {
		goto L62
	}
L34:
	;
	v82 = F_superuser_arg(m, l5)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L45
	}
L35:
	;
	v80 = F_pg_largeobject_aclmask_snapshot(m, l4, l5, v58, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L44
	}
L36:
	;
	v77 = F_object_aclmask_ext(m, int32(2612), l4, l5, v58, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L43
	}
L37:
	;
	v73 = F_object_aclmask_ext(m, int32(1255), l4, l5, v58, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L42
	}
L38:
	;
	v69 = F_object_aclmask_ext(m, int32(1262), l4, l5, v58, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L41
	}
L39:
	;
	v65 = F_pg_class_aclmask_ext(m, l4, l5, v58, int32(1), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	v192 = v65
	goto L24
L41:
	;
	v192 = v69
	goto L24
L42:
	;
	v192 = v73
	goto L24
L43:
	;
	v192 = v77
	goto L24
L44:
	;
	v192 = v80
	goto L24
L45:
	;
	if v82 != 0 {
		v192 = v58
		goto L24
	} else {
		goto L46
	}
L46:
	;
	v85 = F_SearchSysCache1(m, int32(44), l4)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v85 == int32(0) {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v93 = F_SysCacheGetAttr(m, int32(44), v85, int32(3), v16+int32(191))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+191)))
	if v95 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v109 = F_aclmask(m, v106, l5, int32(10), v58, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L56
	}
L51:
	;
	v101 = F_acldefault(m, int32(27), int32(10))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v103 = F_pg_detoast_datum(m, v93)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L55
	}
L54:
	;
	v105 = int32(0)
	v106 = v101
	goto L50
L55:
	;
	v105 = v93
	v106 = v103
	goto L50
L56:
	;
	v111 = int32(0)
	if base.B2i32(v106 == v111)|base.B2i32(v105 == v106) == v111 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v106)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L13
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_ReleaseCatCache(m, v85)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v192 = v109
	goto L24
L62:
	;
	v192 = v123
	goto L24
L63:
	;
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_15), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(3012), int32(_a_F_restrict_and_check_grant_13))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v192 = v140
	goto L24
L67:
	;
	v192 = v144
	goto L24
L68:
	;
	v192 = v148
	goto L24
L69:
	;
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_14), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(3022), int32(_a_F_restrict_and_check_grant_13))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v192 = v165
	goto L24
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l6
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_12), v16+int32(16))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(3029), int32(_a_F_restrict_and_check_grant_13))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v187 = F_pg_attribute_aclmask_ext(m, l4, l8, l5, v58, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v192 = v184 | v187
	goto L24
L78:
	;
	if l9 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v199 = base.B2i32(l6 == int32(6))
	goto L81
L80:
	;
	v199 = int32(0)
	goto L81
L81:
	;
	if v199 != 0 {
		goto L20
	} else {
		goto L82
	}
L82:
	;
	F_aclcheck_error(m, int32(1), l6, l7)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	goto L22
L84:
	;
	m.G0 = v16 + int32(192)
	return v209
L85:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), v364, int32(_a_F_restrict_and_check_grant_2))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L13
	} else {
		goto L137
	}
L86:
	;
	if v209 == int64(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if v209 == int64(0) {
		goto L113
	} else {
		goto L114
	}
L89:
	;
	v215 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L13
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if l2|base.B2i32(v209 == l3) != 0 {
		goto L84
	} else {
		goto L102
	}
L92:
	;
	v217 = int32(0)
	if base.B2i32(l9 == v217)|base.B2i32(l6 != int32(6)) == v217 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v215 == int32(0) {
		goto L84
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v215 == int32(0) {
		goto L84
	} else {
		goto L99
	}
L96:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_0), v16+int32(48))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v364 = int32(334)
	goto L85
L99:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l7
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_3), v16-int32(-64))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	v364 = int32(339)
	goto L85
L102:
	;
	v253 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	v255 = int32(0)
	if base.B2i32(l9 == v255)|base.B2i32(l6 != int32(6)) == v255 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if v253 == int32(0) {
		goto L84
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v253 == int32(0) {
		goto L84
	} else {
		goto L110
	}
L107:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_4), v16+int32(80))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	v364 = int32(347)
	goto L85
L110:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = l7
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_5), v16+int32(96))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	v364 = int32(352)
	goto L85
L113:
	;
	v291 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L13
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if l2|base.B2i32(v209 == l3) != 0 {
		goto L84
	} else {
		goto L126
	}
L116:
	;
	v293 = int32(0)
	if base.B2i32(l9 == v293)|base.B2i32(l6 != int32(6)) == v293 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if v291 == int32(0) {
		goto L84
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v291 == int32(0) {
		goto L84
	} else {
		goto L123
	}
L120:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_6), v16+int32(112))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	v364 = int32(363)
	goto L85
L123:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = l7
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_7), v16+int32(128))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L13
	} else {
		goto L125
	}
L125:
	;
	v364 = int32(368)
	goto L85
L126:
	;
	v329 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	v331 = int32(0)
	if base.B2i32(l9 == v331)|base.B2i32(l6 != int32(6)) == v331 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if v329 == int32(0) {
		goto L84
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if v329 == int32(0) {
		goto L84
	} else {
		goto L134
	}
L131:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+148)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_8), v16+int32(144))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L13
	} else {
		goto L133
	}
L133:
	;
	v364 = int32(376)
	goto L85
L134:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L13
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = l7
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_9), v16+int32(160))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L13
	} else {
		goto L136
	}
L136:
	;
	v364 = int32(381)
	goto L85
L137:
	;
	goto L84
L138:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = l4
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_16), v16+int32(176))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(3495), int32(_a_F_restrict_and_check_grant_17))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_10), v16+int32(32))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L13
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(2956), int32(_a_F_restrict_and_check_grant_11))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_restriction_is_always_true(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	if v6 != 0 {
		v79 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v79
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v7 != 0 {
		v79 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v9 == int32(52) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v12 != int32(1) {
		v79 = v3
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	goto L21
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	if v15 != 0 {
		v79 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v17 != int32(6) {
		v79 = v3
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v20 != 0 {
		v79 = v3
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+8)))
	if v21 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(1)
L12:
	;
	goto L13
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v27 = F_find_base_rel(m, l0, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+8)))
	if int32(0) < v32 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v36 = F_bms_is_member(m, v32, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	return int32(0)
L19:
	;
	if v36 != 0 {
		v79 = int32(1)
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if base.B2i32(v40 != int32(0)) == int32(0) {
		v79 = v3
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v46 == int32(0) {
		v79 = v3
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 <= int32(0) {
		v79 = v3
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v54 = int32(0)
	goto L25
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v54<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 != int32(318) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v79 = v3
	goto L1
L27:
	;
	v73 = v54 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v73 < v74 {
		v54 = v73
		goto L25
	} else {
		goto L31
	}
L28:
	;
	v66 = F_restriction_is_always_true(m, l0, v62)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	if v66 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	return int32(1)
L31:
	;
	goto L26
}
func F_rmdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_rmdir(m, l0)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v2) {
		*(*int32)(unsafe.Add(mBase, _c_F_rmdir[0])) = int32(0) - v2
		v10 = int32(-1)
	} else {
		v10 = v2
	}
	return v10
}
func F_romanian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v650 int32
	_ = v650
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v688 int32
	_ = v688
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v769 int32
	_ = v769
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v888 int32
	_ = v888
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v932 int32
	_ = v932
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1009 int32
	_ = v1009
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1035 int32
	_ = v1035
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1127 int32
	_ = v1127
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1161 int32
	_ = v1161
	var v1172 int32
	_ = v1172
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1250 int32
	_ = v1250
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1276 int32
	_ = v1276
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1368 int32
	_ = v1368
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1436 int32
	_ = v1436
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1487 int32
	_ = v1487
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1520 int32
	_ = v1520
	var v1531 int32
	_ = v1531
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1608 int32
	_ = v1608
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1634 int32
	_ = v1634
	var v1646 int32
	_ = v1646
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1743 int32
	_ = v1743
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1794 int32
	_ = v1794
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1820 int32
	_ = v1820
	var v1827 int32
	_ = v1827
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1916 int32
	_ = v1916
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1942 int32
	_ = v1942
	var v1950 int32
	_ = v1950
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1990 int32
	_ = v1990
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2012 int32
	_ = v2012
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2041 int32
	_ = v2041
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2067 int32
	_ = v2067
	var v2074 int32
	_ = v2074
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2112 int32
	_ = v2112
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	var v2163 int32
	_ = v2163
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2189 int32
	_ = v2189
	var v2197 int32
	_ = v2197
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2275 int32
	_ = v2275
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2330 int32
	_ = v2330
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2378 int32
	_ = v2378
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2450 int32
	_ = v2450
	var v2454 int32
	_ = v2454
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2523 int32
	_ = v2523
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2546 int32
	_ = v2546
	var v2564 int32
	_ = v2564
	var v2566 int32
	_ = v2566
	var v2574 int32
	_ = v2574
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2586 int32
	_ = v2586
	var v2603 int32
	_ = v2603
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2621 int32
	_ = v2621
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2638 int32
	_ = v2638
	var v2641 int32
	_ = v2641
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2652 int32
	_ = v2652
	var v2658 int32
	_ = v2658
	var v2660 int32
	_ = v2660
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2669 int32
	_ = v2669
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2684 int32
	_ = v2684
	var v2692 int32
	_ = v2692
	var v2694 int32
	_ = v2694
	var v2696 int32
	_ = v2696
	var v2698 int32
	_ = v2698
	var v2700 int32
	_ = v2700
	var v2701 int32
	_ = v2701
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2716 int32
	_ = v2716
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2725 int32
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2740 int32
	_ = v2740
	var v2742 int32
	_ = v2742
	var v2747 int32
	_ = v2747
	var v2749 int32
	_ = v2749
	var v2756 int32
	_ = v2756
	var v2759 int32
	_ = v2759
	var v2763 int32
	_ = v2763
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2785 int32
	_ = v2785
	var v2791 int32
	_ = v2791
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2805 int32
	_ = v2805
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v7
	goto L2
L1:
	;
	return v2805
L2:
	;
	v15 = v10 + int32(1)
	goto L4
L3:
	;
	v116 = v7
	goto L43
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 <= v15 {
		v38 = v23
		v39 = v24
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	goto L4
L8:
	;
	v108 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L14
	} else {
		goto L41
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
	v10 = v92
	goto L2
L10:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v95
	switch v32 - int32(1) {
	case 0:
		goto L8
	case 1:
		goto L38
	default:
		goto L7
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	goto L19
L12:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v15))))
	switch v27 - int32(159) {
	case 0, 4:
		goto L13
	default:
		v38 = v23
		v39 = v24
		goto L11
	}
L13:
	;
	v32 = F_find_among(m, l0, int32(_a_F_romanian_UTF_8_stem_1), int32(2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v32 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = v37
	v39 = v36
	goto L11
L17:
	;
	if int32(0) <= v92 {
		goto L9
	} else {
		goto L37
	}
L19:
	;
	goto L20
L20:
	;
	goto L21
L21:
	;
	v47 = v10
	v49 = int32(1)
	goto L24
L23:
	;
	v92 = v77
	goto L17
L24:
	;
	if v39 <= v47 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	v92 = int32(-1)
	goto L17
L27:
	;
	goto L28
L28:
	;
	v54 = v47 + int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v47))))
	if base.Ui32(v56) < base.Ui32(int32(192)) {
		v77 = v54
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v78 = int32(1)
	if v78 < v49 {
		v47 = v77
		v49 = v49 - v78
		goto L24
	} else {
		goto L36
	}
L30:
	;
	if v39 <= v54 {
		v77 = v54
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v63 = v54
	goto L32
L32:
	;
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38+v63))))
	if int32(-65) < v66 {
		v77 = v63
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v77 = v39
	goto L29
L34:
	;
	v70 = v63 + int32(1)
	if v70 != v39 {
		v63 = v70
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L25
L37:
	;
	goto L6
L38:
	;
	v101 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_2))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	if int32(0) <= v101 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v2805 = v101
	goto L1
L41:
	;
	if v108 < int32(0) {
		v2805 = v108
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L7
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L47
L44:
	;
	v2805 = v2800
	goto L1
L45:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v238 != 0 {
		v512 = v239
		goto L70
	} else {
		goto L71
	}
L46:
	;
	v238 = v231
	goto L45
L47:
	;
	if v133 <= v116 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v231 = int32(0)
	goto L46
L49:
	;
	v238 = int32(-1)
	goto L45
L50:
	;
	goto L51
L51:
	;
	v149 = int32(1)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v134))))
	if base.Ui32(v151) < base.Ui32(int32(192)) {
		v208 = v151
		v209 = v149
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if int32(259) < v208 {
		v231 = v209
		goto L46
	} else {
		goto L65
	}
L53:
	;
	v155 = v116 + int32(1)
	if v155 == v133 {
		v208 = v151
		v209 = v149
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v134))))
	v160 = v158 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v151) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v134))))
	v176 = v174 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v151) {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v164 = v116 + int32(2)
	if v164 != v133 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v208 = v151<<(uint(int32(6))%32)&int32(1984) | v160
	v209 = int32(2)
	goto L52
L59:
	;
	goto L58
L60:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v180))))
	v208 = v193&int32(63) | (v151<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v160<<(uint(int32(12))%32) | v176<<(uint(int32(6))%32))
	v209 = int32(4)
	goto L52
L61:
	;
	v180 = v116 + int32(3)
	if v180 != v133 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v208 = v151<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v160<<(uint(int32(6))%32) | v176
	v209 = int32(3)
	goto L52
L64:
	;
	goto L63
L65:
	;
	v213 = v208 - int32(97)
	if v213 < int32(0) {
		v231 = v209
		goto L46
	} else {
		goto L66
	}
L66:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v213)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v219)>>(uint(v213&int32(7))%32))&int32(1) == int32(0) {
		v231 = v209
		goto L46
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v209 + v116
	goto L68
L68:
	;
	goto L48
L69:
	;
	v2800 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_5))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L14
	} else {
		goto L687
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L135
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v240
	if v239 == v240 {
		v379 = v239
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v240
	if v379 == v240 {
		goto L104
	} else {
		goto L105
	}
L73:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v240))))
	if v245 != int32(117) {
		v379 = v239
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v249 = v240 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L77
L75:
	;
	if v369 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L76:
	;
	v369 = v362
	goto L75
L77:
	;
	if v264 <= v249 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v362 = int32(0)
	goto L76
L79:
	;
	v369 = int32(-1)
	goto L75
L80:
	;
	goto L81
L81:
	;
	v280 = int32(1)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249+v265))))
	if base.Ui32(v282) < base.Ui32(int32(192)) {
		v339 = v282
		v340 = v280
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if int32(259) < v339 {
		v362 = v340
		goto L76
	} else {
		goto L95
	}
L83:
	;
	v286 = v240 + int32(2)
	if v286 == v264 {
		v339 = v282
		v340 = v280
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v265))))
	v291 = v289 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v282) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v265))))
	v307 = v305 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v282) {
		goto L91
	} else {
		goto L92
	}
L86:
	;
	v295 = v240 + int32(3)
	if v295 != v264 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v339 = v282<<(uint(int32(6))%32)&int32(1984) | v291
	v340 = int32(2)
	goto L82
L89:
	;
	goto L88
L90:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v311))))
	v339 = v324&int32(63) | (v282<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v291<<(uint(int32(12))%32) | v307<<(uint(int32(6))%32))
	v340 = int32(4)
	goto L82
L91:
	;
	v311 = v240 + int32(4)
	if v311 != v264 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v339 = v282<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v291<<(uint(int32(6))%32) | v307
	v340 = int32(3)
	goto L82
L94:
	;
	goto L93
L95:
	;
	v344 = v339 - int32(97)
	if v344 < int32(0) {
		v362 = v340
		goto L76
	} else {
		goto L96
	}
L96:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v344)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v350)>>(uint(v344&int32(7))%32))&int32(1) == int32(0) {
		v362 = v340
		goto L76
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v340 + v249
	goto L98
L98:
	;
	goto L78
L99:
	;
	v374 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_6))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L14
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v379 = v378
	goto L72
L102:
	;
	if v374 < int32(0) {
		v2805 = v374
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L43
L104:
	;
	v512 = v240
	goto L70
L105:
	;
	goto L106
L106:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382+v240))))
	if v384 != int32(105) {
		v512 = v379
		goto L70
	} else {
		goto L107
	}
L107:
	;
	v388 = v240 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v388
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L110
L108:
	;
	if v508 == int32(0) {
		goto L69
	} else {
		goto L132
	}
L109:
	;
	v508 = v501
	goto L108
L110:
	;
	if v403 <= v388 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v501 = int32(0)
	goto L109
L112:
	;
	v508 = int32(-1)
	goto L108
L113:
	;
	goto L114
L114:
	;
	v419 = int32(1)
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388+v404))))
	if base.Ui32(v421) < base.Ui32(int32(192)) {
		v478 = v421
		v479 = v419
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if int32(259) < v478 {
		v501 = v479
		goto L109
	} else {
		goto L128
	}
L116:
	;
	v425 = v240 + int32(2)
	if v425 == v403 {
		v478 = v421
		v479 = v419
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+v404))))
	v430 = v428 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v421) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434+v404))))
	v446 = v444 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v421) {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	v434 = v240 + int32(3)
	if v434 != v403 {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v478 = v421<<(uint(int32(6))%32)&int32(1984) | v430
	v479 = int32(2)
	goto L115
L122:
	;
	goto L121
L123:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v450))))
	v478 = v463&int32(63) | (v421<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v430<<(uint(int32(12))%32) | v446<<(uint(int32(6))%32))
	v479 = int32(4)
	goto L115
L124:
	;
	v450 = v240 + int32(4)
	if v450 != v403 {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v478 = v421<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v430<<(uint(int32(6))%32) | v446
	v479 = int32(3)
	goto L115
L127:
	;
	goto L126
L128:
	;
	v483 = v478 - int32(97)
	if v483 < int32(0) {
		v501 = v479
		goto L109
	} else {
		goto L129
	}
L129:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v483)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v489)>>(uint(v483&int32(7))%32))&int32(1) == int32(0) {
		v501 = v479
		goto L109
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v479 + v388
	goto L131
L131:
	;
	goto L111
L132:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v512 = v511
	goto L70
L133:
	;
	if int32(0) <= v567 {
		v116 = v567
		goto L43
	} else {
		goto L153
	}
L135:
	;
	goto L136
L136:
	;
	goto L137
L137:
	;
	v522 = v116
	v524 = int32(1)
	goto L140
L139:
	;
	v567 = v552
	goto L133
L140:
	;
	if v512 <= v522 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	v567 = int32(-1)
	goto L133
L143:
	;
	goto L144
L144:
	;
	v529 = v522 + int32(1)
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515+v522))))
	if base.Ui32(v531) < base.Ui32(int32(192)) {
		v552 = v529
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v553 = int32(1)
	if v553 < v524 {
		v522 = v552
		v524 = v524 - v553
		goto L140
	} else {
		goto L152
	}
L146:
	;
	if v512 <= v529 {
		v552 = v529
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v538 = v529
	goto L148
L148:
	;
	v541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v515+v538))))
	if int32(-65) < v541 {
		v552 = v538
		goto L145
	} else {
		goto L150
	}
L149:
	;
	v552 = v512
	goto L145
L150:
	;
	v545 = v538 + int32(1)
	if v545 != v512 {
		v538 = v545
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	goto L141
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v571)+8)) = v572
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v575
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L160
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v577
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1743 = v577
	goto L416
L155:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1717)+8)) = v1715
	goto L154
L156:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1715 = v1713 + v1711
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v577
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L289
L158:
	;
	if v695 != 0 {
		goto L157
	} else {
		goto L182
	}
L159:
	;
	v695 = v688
	goto L158
L160:
	;
	if v590 <= v577 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v688 = int32(0)
	goto L159
L162:
	;
	v695 = int32(-1)
	goto L158
L163:
	;
	goto L164
L164:
	;
	v606 = int32(1)
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577+v591))))
	if base.Ui32(v608) < base.Ui32(int32(192)) {
		v665 = v608
		v666 = v606
		goto L165
	} else {
		goto L166
	}
L165:
	;
	if int32(259) < v665 {
		v688 = v666
		goto L159
	} else {
		goto L178
	}
L166:
	;
	v612 = v577 + int32(1)
	if v612 == v590 {
		v665 = v608
		v666 = v606
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612+v591))))
	v617 = v615 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v608) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v591))))
	v633 = v631 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v608) {
		goto L174
	} else {
		goto L175
	}
L169:
	;
	v621 = v577 + int32(2)
	if v621 != v590 {
		goto L168
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v665 = v608<<(uint(int32(6))%32)&int32(1984) | v617
	v666 = int32(2)
	goto L165
L172:
	;
	goto L171
L173:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591+v637))))
	v665 = v650&int32(63) | (v608<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v617<<(uint(int32(12))%32) | v633<<(uint(int32(6))%32))
	v666 = int32(4)
	goto L165
L174:
	;
	v637 = v577 + int32(3)
	if v637 != v590 {
		goto L173
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v665 = v608<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v617<<(uint(int32(6))%32) | v633
	v666 = int32(3)
	goto L165
L177:
	;
	goto L176
L178:
	;
	v670 = v665 - int32(97)
	if v670 < int32(0) {
		v688 = v666
		goto L159
	} else {
		goto L179
	}
L179:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v670)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v676)>>(uint(v670&int32(7))%32))&int32(1) == int32(0) {
		v688 = v666
		goto L159
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v666 + v577
	goto L181
L181:
	;
	goto L161
L182:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L185
L183:
	;
	if v813 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L184:
	;
	v813 = v806
	goto L183
L185:
	;
	if v709 <= v696 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v806 = int32(0)
	goto L184
L187:
	;
	v813 = int32(-1)
	goto L183
L188:
	;
	goto L189
L189:
	;
	v725 = int32(1)
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v710))))
	if base.Ui32(v727) < base.Ui32(int32(192)) {
		v784 = v727
		v785 = v725
		goto L190
	} else {
		goto L191
	}
L190:
	;
	if int32(259) < v784 {
		goto L203
	} else {
		goto L204
	}
L191:
	;
	v731 = v696 + int32(1)
	if v731 == v709 {
		v784 = v727
		v785 = v725
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+v710))))
	v736 = v734 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v727) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740+v710))))
	v752 = v750 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v727) {
		goto L199
	} else {
		goto L200
	}
L194:
	;
	v740 = v696 + int32(2)
	if v740 != v709 {
		goto L193
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v784 = v727<<(uint(int32(6))%32)&int32(1984) | v736
	v785 = int32(2)
	goto L190
L197:
	;
	goto L196
L198:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710+v756))))
	v784 = v769&int32(63) | (v727<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v736<<(uint(int32(12))%32) | v752<<(uint(int32(6))%32))
	v785 = int32(4)
	goto L190
L199:
	;
	v756 = v696 + int32(3)
	if v756 != v709 {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v784 = v727<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v736<<(uint(int32(6))%32) | v752
	v785 = int32(3)
	goto L190
L202:
	;
	goto L201
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v785 + v696
	goto L207
L204:
	;
	v789 = v784 - int32(97)
	if v789 < int32(0) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v789)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v795)>>(uint(v789&int32(7))%32))&int32(1) != 0 {
		v806 = v785
		goto L184
	} else {
		goto L206
	}
L206:
	;
	goto L203
L207:
	;
	goto L186
L208:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v837 = v827
	goto L213
L209:
	;
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v696
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L239
L211:
	;
	if int32(0) <= v932 {
		v1711 = v932
		goto L156
	} else {
		goto L236
	}
L212:
	;
	v932 = v904
	goto L211
L213:
	;
	if v828 <= v837 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v932 = int32(-1)
	goto L211
L216:
	;
	goto L217
L217:
	;
	v844 = int32(1)
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837+v829))))
	if base.Ui32(v846) < base.Ui32(int32(192)) {
		v903 = v846
		v904 = v844
		goto L218
	} else {
		goto L219
	}
L218:
	;
	if int32(259) < v903 {
		goto L231
	} else {
		goto L232
	}
L219:
	;
	v850 = v837 + int32(1)
	if v850 == v828 {
		v903 = v846
		v904 = v844
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850+v829))))
	v855 = v853 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v846) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859+v829))))
	v871 = v869 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v846) {
		goto L227
	} else {
		goto L228
	}
L222:
	;
	v859 = v837 + int32(2)
	if v859 != v828 {
		goto L221
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v903 = v846<<(uint(int32(6))%32)&int32(1984) | v855
	v904 = int32(2)
	goto L218
L225:
	;
	goto L224
L226:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829+v875))))
	v903 = v888&int32(63) | (v846<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v855<<(uint(int32(12))%32) | v871<<(uint(int32(6))%32))
	v904 = int32(4)
	goto L218
L227:
	;
	v875 = v837 + int32(3)
	if v875 != v828 {
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v903 = v846<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v855<<(uint(int32(6))%32) | v871
	v904 = int32(3)
	goto L218
L230:
	;
	goto L229
L231:
	;
	v921 = v904 + v837
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v921
	v837 = v921
	goto L213
L232:
	;
	v908 = v903 - int32(97)
	if v908 < int32(0) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v908)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v914)>>(uint(v908&int32(7))%32))&int32(1) != 0 {
		goto L212
	} else {
		goto L234
	}
L234:
	;
	goto L231
L236:
	;
	goto L210
L237:
	;
	if v1054 != 0 {
		goto L157
	} else {
		goto L261
	}
L238:
	;
	v1054 = v1047
	goto L237
L239:
	;
	if v949 <= v696 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1047 = int32(0)
	goto L238
L241:
	;
	v1054 = int32(-1)
	goto L237
L242:
	;
	goto L243
L243:
	;
	v965 = int32(1)
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v950))))
	if base.Ui32(v967) < base.Ui32(int32(192)) {
		v1024 = v967
		v1025 = v965
		goto L244
	} else {
		goto L245
	}
L244:
	;
	if int32(259) < v1024 {
		v1047 = v1025
		goto L238
	} else {
		goto L257
	}
L245:
	;
	v971 = v696 + int32(1)
	if v971 == v949 {
		v1024 = v967
		v1025 = v965
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971+v950))))
	v976 = v974 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v967) {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980+v950))))
	v992 = v990 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v967) {
		goto L253
	} else {
		goto L254
	}
L248:
	;
	v980 = v696 + int32(2)
	if v980 != v949 {
		goto L247
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1024 = v967<<(uint(int32(6))%32)&int32(1984) | v976
	v1025 = int32(2)
	goto L244
L251:
	;
	goto L250
L252:
	;
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950+v996))))
	v1024 = v1009&int32(63) | (v967<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v976<<(uint(int32(12))%32) | v992<<(uint(int32(6))%32))
	v1025 = int32(4)
	goto L244
L253:
	;
	v996 = v696 + int32(3)
	if v996 != v949 {
		goto L252
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1024 = v967<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v976<<(uint(int32(6))%32) | v992
	v1025 = int32(3)
	goto L244
L256:
	;
	goto L255
L257:
	;
	v1029 = v1024 - int32(97)
	if v1029 < int32(0) {
		v1047 = v1025
		goto L238
	} else {
		goto L258
	}
L258:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1029)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1035)>>(uint(v1029&int32(7))%32))&int32(1) == int32(0) {
		v1047 = v1025
		goto L238
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1025 + v696
	goto L260
L260:
	;
	goto L240
L261:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1076 = v1066
	goto L264
L262:
	;
	if int32(0) <= v1172 {
		v1711 = v1172
		goto L156
	} else {
		goto L286
	}
L263:
	;
	v1172 = v1143
	goto L262
L264:
	;
	if v1067 <= v1076 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1172 = int32(-1)
	goto L262
L267:
	;
	goto L268
L268:
	;
	v1083 = int32(1)
	v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076+v1068))))
	if base.Ui32(v1085) < base.Ui32(int32(192)) {
		v1142 = v1085
		v1143 = v1083
		goto L269
	} else {
		goto L270
	}
L269:
	;
	if int32(259) < v1142 {
		goto L263
	} else {
		goto L282
	}
L270:
	;
	v1089 = v1076 + int32(1)
	if v1089 == v1067 {
		v1142 = v1085
		v1143 = v1083
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089+v1068))))
	v1094 = v1092 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1085) {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1098+v1068))))
	v1110 = v1108 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1085) {
		goto L278
	} else {
		goto L279
	}
L273:
	;
	v1098 = v1076 + int32(2)
	if v1098 != v1067 {
		goto L272
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1142 = v1085<<(uint(int32(6))%32)&int32(1984) | v1094
	v1143 = int32(2)
	goto L269
L276:
	;
	goto L275
L277:
	;
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068+v1114))))
	v1142 = v1127&int32(63) | (v1085<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v1094<<(uint(int32(12))%32) | v1110<<(uint(int32(6))%32))
	v1143 = int32(4)
	goto L269
L278:
	;
	v1114 = v1076 + int32(3)
	if v1114 != v1067 {
		goto L277
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1142 = v1085<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1094<<(uint(int32(6))%32) | v1110
	v1143 = int32(3)
	goto L269
L281:
	;
	goto L280
L282:
	;
	v1147 = v1142 - int32(97)
	if v1147 < int32(0) {
		goto L263
	} else {
		goto L283
	}
L283:
	;
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1147)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1153)>>(uint(v1147&int32(7))%32))&int32(1) == int32(0) {
		goto L263
	} else {
		goto L284
	}
L284:
	;
	v1161 = v1143 + v1076
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1161
	v1076 = v1161
	goto L264
L286:
	;
	goto L157
L287:
	;
	if v1294 != 0 {
		goto L154
	} else {
		goto L312
	}
L288:
	;
	v1294 = v1287
	goto L287
L289:
	;
	if v1190 <= v577 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1287 = int32(0)
	goto L288
L291:
	;
	v1294 = int32(-1)
	goto L287
L292:
	;
	goto L293
L293:
	;
	v1206 = int32(1)
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577+v1191))))
	if base.Ui32(v1208) < base.Ui32(int32(192)) {
		v1265 = v1208
		v1266 = v1206
		goto L294
	} else {
		goto L295
	}
L294:
	;
	if int32(259) < v1265 {
		goto L307
	} else {
		goto L308
	}
L295:
	;
	v1212 = v577 + int32(1)
	if v1212 == v1190 {
		v1265 = v1208
		v1266 = v1206
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1212+v1191))))
	v1217 = v1215 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1208) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221+v1191))))
	v1233 = v1231 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1208) {
		goto L303
	} else {
		goto L304
	}
L298:
	;
	v1221 = v577 + int32(2)
	if v1221 != v1190 {
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1265 = v1208<<(uint(int32(6))%32)&int32(1984) | v1217
	v1266 = int32(2)
	goto L294
L301:
	;
	goto L300
L302:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191+v1237))))
	v1265 = v1250&int32(63) | (v1208<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v1217<<(uint(int32(12))%32) | v1233<<(uint(int32(6))%32))
	v1266 = int32(4)
	goto L294
L303:
	;
	v1237 = v577 + int32(3)
	if v1237 != v1190 {
		goto L302
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1265 = v1208<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1217<<(uint(int32(6))%32) | v1233
	v1266 = int32(3)
	goto L294
L306:
	;
	goto L305
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1266 + v577
	goto L311
L308:
	;
	v1270 = v1265 - int32(97)
	if v1270 < int32(0) {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1270)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1276)>>(uint(v1270&int32(7))%32))&int32(1) != 0 {
		v1287 = v1266
		goto L288
	} else {
		goto L310
	}
L310:
	;
	goto L307
L311:
	;
	goto L290
L312:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L315
L313:
	;
	if v1412 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L314:
	;
	v1412 = v1405
	goto L313
L315:
	;
	if v1308 <= v1295 {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	v1405 = int32(0)
	goto L314
L317:
	;
	v1412 = int32(-1)
	goto L313
L318:
	;
	goto L319
L319:
	;
	v1324 = int32(1)
	v1326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295+v1309))))
	if base.Ui32(v1326) < base.Ui32(int32(192)) {
		v1383 = v1326
		v1384 = v1324
		goto L320
	} else {
		goto L321
	}
L320:
	;
	if int32(259) < v1383 {
		goto L333
	} else {
		goto L334
	}
L321:
	;
	v1330 = v1295 + int32(1)
	if v1330 == v1308 {
		v1383 = v1326
		v1384 = v1324
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330+v1309))))
	v1335 = v1333 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1326) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1339+v1309))))
	v1351 = v1349 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1326) {
		goto L329
	} else {
		goto L330
	}
L324:
	;
	v1339 = v1295 + int32(2)
	if v1339 != v1308 {
		goto L323
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1383 = v1326<<(uint(int32(6))%32)&int32(1984) | v1335
	v1384 = int32(2)
	goto L320
L327:
	;
	goto L326
L328:
	;
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309+v1355))))
	v1383 = v1368&int32(63) | (v1326<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v1335<<(uint(int32(12))%32) | v1351<<(uint(int32(6))%32))
	v1384 = int32(4)
	goto L320
L329:
	;
	v1355 = v1295 + int32(3)
	if v1355 != v1308 {
		goto L328
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	v1383 = v1326<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1335<<(uint(int32(6))%32) | v1351
	v1384 = int32(3)
	goto L320
L332:
	;
	goto L331
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1384 + v1295
	goto L337
L334:
	;
	v1388 = v1383 - int32(97)
	if v1388 < int32(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1388)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1394)>>(uint(v1388&int32(7))%32))&int32(1) != 0 {
		v1405 = v1384
		goto L314
	} else {
		goto L336
	}
L336:
	;
	goto L333
L337:
	;
	goto L316
L338:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1436 = v1426
	goto L343
L339:
	;
	goto L340
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1295
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L369
L341:
	;
	if int32(0) <= v1531 {
		v1711 = v1531
		goto L156
	} else {
		goto L366
	}
L342:
	;
	v1531 = v1503
	goto L341
L343:
	;
	if v1427 <= v1436 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1531 = int32(-1)
	goto L341
L346:
	;
	goto L347
L347:
	;
	v1443 = int32(1)
	v1445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436+v1428))))
	if base.Ui32(v1445) < base.Ui32(int32(192)) {
		v1502 = v1445
		v1503 = v1443
		goto L348
	} else {
		goto L349
	}
L348:
	;
	if int32(259) < v1502 {
		goto L361
	} else {
		goto L362
	}
L349:
	;
	v1449 = v1436 + int32(1)
	if v1449 == v1427 {
		v1502 = v1445
		v1503 = v1443
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v1452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1449+v1428))))
	v1454 = v1452 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1445) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1458+v1428))))
	v1470 = v1468 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1445) {
		goto L357
	} else {
		goto L358
	}
L352:
	;
	v1458 = v1436 + int32(2)
	if v1458 != v1427 {
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1502 = v1445<<(uint(int32(6))%32)&int32(1984) | v1454
	v1503 = int32(2)
	goto L348
L355:
	;
	goto L354
L356:
	;
	v1487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1428+v1474))))
	v1502 = v1487&int32(63) | (v1445<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v1454<<(uint(int32(12))%32) | v1470<<(uint(int32(6))%32))
	v1503 = int32(4)
	goto L348
L357:
	;
	v1474 = v1436 + int32(3)
	if v1474 != v1427 {
		goto L356
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1502 = v1445<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1454<<(uint(int32(6))%32) | v1470
	v1503 = int32(3)
	goto L348
L360:
	;
	goto L359
L361:
	;
	v1520 = v1503 + v1436
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1520
	v1436 = v1520
	goto L343
L362:
	;
	v1507 = v1502 - int32(97)
	if v1507 < int32(0) {
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1507)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1513)>>(uint(v1507&int32(7))%32))&int32(1) != 0 {
		goto L342
	} else {
		goto L364
	}
L364:
	;
	goto L361
L366:
	;
	goto L340
L367:
	;
	if v1653 != 0 {
		goto L154
	} else {
		goto L391
	}
L368:
	;
	v1653 = v1646
	goto L367
L369:
	;
	if v1548 <= v1295 {
		goto L371
	} else {
		goto L372
	}
L370:
	;
	v1646 = int32(0)
	goto L368
L371:
	;
	v1653 = int32(-1)
	goto L367
L372:
	;
	goto L373
L373:
	;
	v1564 = int32(1)
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295+v1549))))
	if base.Ui32(v1566) < base.Ui32(int32(192)) {
		v1623 = v1566
		v1624 = v1564
		goto L374
	} else {
		goto L375
	}
L374:
	;
	if int32(259) < v1623 {
		v1646 = v1624
		goto L368
	} else {
		goto L387
	}
L375:
	;
	v1570 = v1295 + int32(1)
	if v1570 == v1548 {
		v1623 = v1566
		v1624 = v1564
		goto L374
	} else {
		goto L376
	}
L376:
	;
	v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1570+v1549))))
	v1575 = v1573 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1566) {
		goto L378
	} else {
		goto L379
	}
L377:
	;
	v1589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1579+v1549))))
	v1591 = v1589 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1566) {
		goto L383
	} else {
		goto L384
	}
L378:
	;
	v1579 = v1295 + int32(2)
	if v1579 != v1548 {
		goto L377
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v1623 = v1566<<(uint(int32(6))%32)&int32(1984) | v1575
	v1624 = int32(2)
	goto L374
L381:
	;
	goto L380
L382:
	;
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549+v1595))))
	v1623 = v1608&int32(63) | (v1566<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v1575<<(uint(int32(12))%32) | v1591<<(uint(int32(6))%32))
	v1624 = int32(4)
	goto L374
L383:
	;
	v1595 = v1295 + int32(3)
	if v1595 != v1548 {
		goto L382
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v1623 = v1566<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1575<<(uint(int32(6))%32) | v1591
	v1624 = int32(3)
	goto L374
L386:
	;
	goto L385
L387:
	;
	v1628 = v1623 - int32(97)
	if v1628 < int32(0) {
		v1646 = v1624
		goto L368
	} else {
		goto L388
	}
L388:
	;
	v1634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1628)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1634)>>(uint(v1628&int32(7))%32))&int32(1) == int32(0) {
		v1646 = v1624
		goto L368
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1624 + v1295
	goto L390
L390:
	;
	goto L370
L391:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L394
L392:
	;
	if int32(0) <= v1708 {
		v1715 = v1708
		goto L155
	} else {
		goto L412
	}
L394:
	;
	goto L395
L395:
	;
	goto L396
L396:
	;
	v1663 = v1655
	v1665 = int32(1)
	goto L399
L398:
	;
	v1708 = v1693
	goto L392
L399:
	;
	if v1656 <= v1663 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	goto L398
L401:
	;
	v1708 = int32(-1)
	goto L392
L402:
	;
	goto L403
L403:
	;
	v1670 = v1663 + int32(1)
	v1672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654+v1663))))
	if base.Ui32(v1672) < base.Ui32(int32(192)) {
		v1693 = v1670
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1694 = int32(1)
	if v1694 < v1665 {
		v1663 = v1693
		v1665 = v1665 - v1694
		goto L399
	} else {
		goto L411
	}
L405:
	;
	if v1656 <= v1670 {
		v1693 = v1670
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v1679 = v1670
	goto L407
L407:
	;
	v1682 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1654+v1679))))
	if int32(-65) < v1682 {
		v1693 = v1679
		goto L404
	} else {
		goto L409
	}
L408:
	;
	v1693 = v1656
	goto L404
L409:
	;
	v1686 = v1679 + int32(1)
	if v1686 != v1656 {
		v1679 = v1686
		goto L407
	} else {
		goto L410
	}
L410:
	;
	goto L408
L411:
	;
	goto L400
L412:
	;
	goto L154
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v577
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2217
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2217
	v2221 = v2217 - int32(1)
	if v2221 <= v577 {
		goto L516
	} else {
		goto L517
	}
L414:
	;
	if v1838 < int32(0) {
		goto L413
	} else {
		goto L439
	}
L415:
	;
	v1838 = v1810
	goto L414
L416:
	;
	if v1734 <= v1743 {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1838 = int32(-1)
	goto L414
L419:
	;
	goto L420
L420:
	;
	v1750 = int32(1)
	v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1743+v1735))))
	if base.Ui32(v1752) < base.Ui32(int32(192)) {
		v1809 = v1752
		v1810 = v1750
		goto L421
	} else {
		goto L422
	}
L421:
	;
	if int32(259) < v1809 {
		goto L434
	} else {
		goto L435
	}
L422:
	;
	v1756 = v1743 + int32(1)
	if v1756 == v1734 {
		v1809 = v1752
		v1810 = v1750
		goto L421
	} else {
		goto L423
	}
L423:
	;
	v1759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1756+v1735))))
	v1761 = v1759 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1752) {
		goto L425
	} else {
		goto L426
	}
L424:
	;
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1765+v1735))))
	v1777 = v1775 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1752) {
		goto L430
	} else {
		goto L431
	}
L425:
	;
	v1765 = v1743 + int32(2)
	if v1765 != v1734 {
		goto L424
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	v1809 = v1752<<(uint(int32(6))%32)&int32(1984) | v1761
	v1810 = int32(2)
	goto L421
L428:
	;
	goto L427
L429:
	;
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1735+v1781))))
	v1809 = v1794&int32(63) | (v1752<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v1761<<(uint(int32(12))%32) | v1777<<(uint(int32(6))%32))
	v1810 = int32(4)
	goto L421
L430:
	;
	v1781 = v1743 + int32(3)
	if v1781 != v1734 {
		goto L429
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	v1809 = v1752<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1761<<(uint(int32(6))%32) | v1777
	v1810 = int32(3)
	goto L421
L433:
	;
	goto L432
L434:
	;
	v1827 = v1810 + v1743
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1827
	v1743 = v1827
	goto L416
L435:
	;
	v1814 = v1809 - int32(97)
	if v1814 < int32(0) {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1814)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1820)>>(uint(v1814&int32(7))%32))&int32(1) != 0 {
		goto L415
	} else {
		goto L437
	}
L437:
	;
	goto L434
L439:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1842 = v1841 + v1838
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1842
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1865 = v1842
	goto L442
L440:
	;
	if v1961 < int32(0) {
		goto L413
	} else {
		goto L464
	}
L441:
	;
	v1961 = v1932
	goto L440
L442:
	;
	if v1856 <= v1865 {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1961 = int32(-1)
	goto L440
L445:
	;
	goto L446
L446:
	;
	v1872 = int32(1)
	v1874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865+v1857))))
	if base.Ui32(v1874) < base.Ui32(int32(192)) {
		v1931 = v1874
		v1932 = v1872
		goto L447
	} else {
		goto L448
	}
L447:
	;
	if int32(259) < v1931 {
		goto L441
	} else {
		goto L460
	}
L448:
	;
	v1878 = v1865 + int32(1)
	if v1878 == v1856 {
		v1931 = v1874
		v1932 = v1872
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v1881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1878+v1857))))
	v1883 = v1881 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1874) {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	v1897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1887+v1857))))
	v1899 = v1897 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1874) {
		goto L456
	} else {
		goto L457
	}
L451:
	;
	v1887 = v1865 + int32(2)
	if v1887 != v1856 {
		goto L450
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v1931 = v1874<<(uint(int32(6))%32)&int32(1984) | v1883
	v1932 = int32(2)
	goto L447
L454:
	;
	goto L453
L455:
	;
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1857+v1903))))
	v1931 = v1916&int32(63) | (v1874<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v1883<<(uint(int32(12))%32) | v1899<<(uint(int32(6))%32))
	v1932 = int32(4)
	goto L447
L456:
	;
	v1903 = v1865 + int32(3)
	if v1903 != v1856 {
		goto L455
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	v1931 = v1874<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1883<<(uint(int32(6))%32) | v1899
	v1932 = int32(3)
	goto L447
L459:
	;
	goto L458
L460:
	;
	v1936 = v1931 - int32(97)
	if v1936 < int32(0) {
		goto L441
	} else {
		goto L461
	}
L461:
	;
	v1942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1936)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1942)>>(uint(v1936&int32(7))%32))&int32(1) == int32(0) {
		goto L441
	} else {
		goto L462
	}
L462:
	;
	v1950 = v1932 + v1865
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1950
	v1865 = v1950
	goto L442
L464:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1965 = v1964 + v1961
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1965
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1967)+4)) = v1965
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1990 = v1980
	goto L467
L465:
	;
	if v2085 < int32(0) {
		goto L413
	} else {
		goto L490
	}
L466:
	;
	v2085 = v2057
	goto L465
L467:
	;
	if v1981 <= v1990 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v2085 = int32(-1)
	goto L465
L470:
	;
	goto L471
L471:
	;
	v1997 = int32(1)
	v1999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1990+v1982))))
	if base.Ui32(v1999) < base.Ui32(int32(192)) {
		v2056 = v1999
		v2057 = v1997
		goto L472
	} else {
		goto L473
	}
L472:
	;
	if int32(259) < v2056 {
		goto L485
	} else {
		goto L486
	}
L473:
	;
	v2003 = v1990 + int32(1)
	if v2003 == v1981 {
		v2056 = v1999
		v2057 = v1997
		goto L472
	} else {
		goto L474
	}
L474:
	;
	v2006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2003+v1982))))
	v2008 = v2006 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1999) {
		goto L476
	} else {
		goto L477
	}
L475:
	;
	v2022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2012+v1982))))
	v2024 = v2022 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1999) {
		goto L481
	} else {
		goto L482
	}
L476:
	;
	v2012 = v1990 + int32(2)
	if v2012 != v1981 {
		goto L475
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v2056 = v1999<<(uint(int32(6))%32)&int32(1984) | v2008
	v2057 = int32(2)
	goto L472
L479:
	;
	goto L478
L480:
	;
	v2041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1982+v2028))))
	v2056 = v2041&int32(63) | (v1999<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v2008<<(uint(int32(12))%32) | v2024<<(uint(int32(6))%32))
	v2057 = int32(4)
	goto L472
L481:
	;
	v2028 = v1990 + int32(3)
	if v2028 != v1981 {
		goto L480
	} else {
		goto L484
	}
L482:
	;
	goto L483
L483:
	;
	v2056 = v1999<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v2008<<(uint(int32(6))%32) | v2024
	v2057 = int32(3)
	goto L472
L484:
	;
	goto L483
L485:
	;
	v2074 = v2057 + v1990
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2074
	v1990 = v2074
	goto L467
L486:
	;
	v2061 = v2056 - int32(97)
	if v2061 < int32(0) {
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v2067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2061)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2067)>>(uint(v2061&int32(7))%32))&int32(1) != 0 {
		goto L466
	} else {
		goto L488
	}
L488:
	;
	goto L485
L490:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2089 = v2088 + v2085
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2089
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2112 = v2089
	goto L493
L491:
	;
	if v2208 < int32(0) {
		goto L413
	} else {
		goto L515
	}
L492:
	;
	v2208 = v2179
	goto L491
L493:
	;
	if v2103 <= v2112 {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	v2208 = int32(-1)
	goto L491
L496:
	;
	goto L497
L497:
	;
	v2119 = int32(1)
	v2121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2112+v2104))))
	if base.Ui32(v2121) < base.Ui32(int32(192)) {
		v2178 = v2121
		v2179 = v2119
		goto L498
	} else {
		goto L499
	}
L498:
	;
	if int32(259) < v2178 {
		goto L492
	} else {
		goto L511
	}
L499:
	;
	v2125 = v2112 + int32(1)
	if v2125 == v2103 {
		v2178 = v2121
		v2179 = v2119
		goto L498
	} else {
		goto L500
	}
L500:
	;
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2125+v2104))))
	v2130 = v2128 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2121) {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	v2144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2134+v2104))))
	v2146 = v2144 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2121) {
		goto L507
	} else {
		goto L508
	}
L502:
	;
	v2134 = v2112 + int32(2)
	if v2134 != v2103 {
		goto L501
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v2178 = v2121<<(uint(int32(6))%32)&int32(1984) | v2130
	v2179 = int32(2)
	goto L498
L505:
	;
	goto L504
L506:
	;
	v2163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2150))))
	v2178 = v2163&int32(63) | (v2121<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_3) | v2130<<(uint(int32(12))%32) | v2146<<(uint(int32(6))%32))
	v2179 = int32(4)
	goto L498
L507:
	;
	v2150 = v2112 + int32(3)
	if v2150 != v2103 {
		goto L506
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	v2178 = v2121<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v2130<<(uint(int32(6))%32) | v2146
	v2179 = int32(3)
	goto L498
L510:
	;
	goto L509
L511:
	;
	v2183 = v2178 - int32(97)
	if v2183 < int32(0) {
		goto L492
	} else {
		goto L512
	}
L512:
	;
	v2189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2183)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2189)>>(uint(v2183&int32(7))%32))&int32(1) == int32(0) {
		goto L492
	} else {
		goto L513
	}
L513:
	;
	v2197 = v2179 + v2112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2197
	v2112 = v2197
	goto L493
L515:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2211))) = v2212 + v2208
	goto L413
L516:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2313
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2315)+12)) = int32(0)
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2318
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2321 = v2318 - v2320
	v2324 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_7), int32(46))
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L14
	} else {
		goto L549
	}
L517:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2223+v2221))))
	if base.B2i32(v2225&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2225)%32)&int32(_a_F_romanian_UTF_8_stem_8) == int32(0)) != 0 {
		goto L516
	} else {
		goto L518
	}
L518:
	;
	v2239 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_9), int32(16))
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L14
	} else {
		goto L519
	}
L519:
	;
	if v2239 == int32(0) {
		goto L516
	} else {
		goto L520
	}
L520:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2243
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+4))
	if v2243 < v2246 {
		goto L516
	} else {
		goto L521
	}
L521:
	;
	switch v2239 - int32(1) {
	case 0:
		goto L528
	case 1:
		goto L527
	case 2:
		goto L526
	case 3:
		goto L525
	case 4:
		goto L524
	case 5:
		goto L523
	case 6:
		goto L522
	default:
		goto L516
	}
L522:
	;
	v2307 = F_slice_from_s(m, l0, int32(4), int32(_a_F_romanian_UTF_8_stem_10))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L14
	} else {
		goto L546
	}
L523:
	;
	v2301 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_11))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L14
	} else {
		goto L544
	}
L524:
	;
	v2272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2273 = int32(2)
	v2275 = int32(0)
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2277-v2278 < v2273 {
		v2288 = v2275
		goto L538
	} else {
		goto L539
	}
L525:
	;
	v2268 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_12))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L14
	} else {
		goto L535
	}
L526:
	;
	v2262 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_13))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L14
	} else {
		goto L533
	}
L527:
	;
	v2256 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_14))
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L14
	} else {
		goto L531
	}
L528:
	;
	v2250 = F_slice_del(m, l0)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L14
	} else {
		goto L529
	}
L529:
	;
	if int32(0) <= v2250 {
		goto L516
	} else {
		goto L530
	}
L530:
	;
	v2805 = v2250
	goto L1
L531:
	;
	if int32(0) <= v2256 {
		goto L516
	} else {
		goto L532
	}
L532:
	;
	v2805 = v2256
	goto L1
L533:
	;
	if int32(0) <= v2262 {
		goto L516
	} else {
		goto L534
	}
L534:
	;
	v2805 = v2262
	goto L1
L535:
	;
	if int32(0) <= v2268 {
		goto L516
	} else {
		goto L536
	}
L536:
	;
	v2805 = v2268
	goto L1
L537:
	;
	if v2288 != 0 {
		goto L516
	} else {
		goto L541
	}
L538:
	;
	goto L537
L539:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2284 = F_memcmp(m, v2281+v2277-v2273, int32(_a_F_romanian_UTF_8_stem_15), v2273)
	mBase = m.M
	if v2284 != 0 {
		v2288 = v2275
		goto L538
	} else {
		goto L540
	}
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2277 - v2273
	v2288 = int32(1)
	goto L538
L541:
	;
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2289 + (v2243 - v2272)
	v2295 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_16))
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L14
	} else {
		goto L542
	}
L542:
	;
	if int32(0) <= v2295 {
		goto L516
	} else {
		goto L543
	}
L543:
	;
	v2805 = v2295
	goto L1
L544:
	;
	if int32(0) <= v2301 {
		goto L516
	} else {
		goto L545
	}
L545:
	;
	v2805 = v2301
	goto L1
L546:
	;
	if v2307 < int32(0) {
		v2805 = v2307
		goto L1
	} else {
		goto L547
	}
L547:
	;
	goto L516
L548:
	;
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2396 = v2395 + v2321
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2396
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2396
	v2401 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_17), int32(62))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L14
	} else {
		goto L575
	}
L549:
	;
	if v2324 == int32(0) {
		goto L548
	} else {
		goto L550
	}
L550:
	;
	v2330 = v2324
	goto L551
L551:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2334
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+4))
	if v2334 < v2337 {
		goto L548
	} else {
		goto L553
	}
L552:
	;
	goto L548
L553:
	;
	switch v2330 - int32(1) {
	case 0:
		goto L560
	case 1:
		goto L559
	case 2:
		goto L558
	case 3:
		goto L557
	case 4:
		goto L556
	case 5:
		goto L555
	default:
		goto L554
	}
L554:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2378)+12)) = int32(1)
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2382 = v2381 + v2321
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2382
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2382
	v2387 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_7), int32(46))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L14
	} else {
		goto L573
	}
L555:
	;
	v2373 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_18))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L14
	} else {
		goto L571
	}
L556:
	;
	v2367 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_19))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L14
	} else {
		goto L569
	}
L557:
	;
	v2361 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_20))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L14
	} else {
		goto L567
	}
L558:
	;
	v2355 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_21))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L14
	} else {
		goto L565
	}
L559:
	;
	v2349 = F_slice_from_s(m, l0, int32(4), int32(_a_F_romanian_UTF_8_stem_22))
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L14
	} else {
		goto L563
	}
L560:
	;
	v2343 = F_slice_from_s(m, l0, int32(4), int32(_a_F_romanian_UTF_8_stem_23))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L14
	} else {
		goto L561
	}
L561:
	;
	if int32(0) <= v2343 {
		goto L554
	} else {
		goto L562
	}
L562:
	;
	v2805 = v2343
	goto L1
L563:
	;
	if int32(0) <= v2349 {
		goto L554
	} else {
		goto L564
	}
L564:
	;
	v2805 = v2349
	goto L1
L565:
	;
	if int32(0) <= v2355 {
		goto L554
	} else {
		goto L566
	}
L566:
	;
	v2805 = v2355
	goto L1
L567:
	;
	if int32(0) <= v2361 {
		goto L554
	} else {
		goto L568
	}
L568:
	;
	v2805 = v2361
	goto L1
L569:
	;
	if int32(0) <= v2367 {
		goto L554
	} else {
		goto L570
	}
L570:
	;
	v2805 = v2367
	goto L1
L571:
	;
	if v2373 < int32(0) {
		v2805 = v2373
		goto L1
	} else {
		goto L572
	}
L572:
	;
	goto L554
L573:
	;
	if v2387 != 0 {
		v2330 = v2387
		goto L551
	} else {
		goto L574
	}
L574:
	;
	goto L552
L575:
	;
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2401 == int32(0) {
		v2454 = v2403
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2456
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v2454)+12))
	if v2458 != 0 {
		goto L596
	} else {
		goto L597
	}
L577:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2406
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v2403)))
	if v2406 < v2408 {
		v2454 = v2403
		goto L576
	} else {
		goto L578
	}
L578:
	;
	switch v2401 - int32(1) {
	case 0:
		goto L582
	case 1:
		goto L581
	case 2:
		goto L580
	default:
		goto L579
	}
L579:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2450)+12)) = int32(1)
	v2454 = v2450
	goto L576
L580:
	;
	v2445 = F_slice_from_s(m, l0, int32(3), int32(_a_F_romanian_UTF_8_stem_24))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L14
	} else {
		goto L594
	}
L581:
	;
	v2416 = int32(2)
	v2418 = int32(0)
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2420-v2421 < v2416 {
		v2431 = v2418
		goto L586
	} else {
		goto L587
	}
L582:
	;
	v2412 = F_slice_del(m, l0)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L14
	} else {
		goto L583
	}
L583:
	;
	if int32(0) <= v2412 {
		goto L579
	} else {
		goto L584
	}
L584:
	;
	v2805 = v2412
	goto L1
L585:
	;
	if v2431 == int32(0) {
		goto L589
	} else {
		goto L590
	}
L586:
	;
	goto L585
L587:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2427 = F_memcmp(m, v2424+v2420-v2416, int32(_a_F_romanian_UTF_8_stem_25), v2416)
	mBase = m.M
	if v2427 != 0 {
		v2431 = v2418
		goto L586
	} else {
		goto L588
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2420 - v2416
	v2431 = int32(1)
	goto L586
L589:
	;
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2454 = v2434
	goto L576
L590:
	;
	goto L591
L591:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2435
	v2439 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_26))
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L14
	} else {
		goto L592
	}
L592:
	;
	if int32(0) <= v2439 {
		goto L579
	} else {
		goto L593
	}
L593:
	;
	v2805 = v2439
	goto L1
L594:
	;
	if v2445 < int32(0) {
		v2805 = v2445
		goto L1
	} else {
		goto L595
	}
L595:
	;
	goto L579
L596:
	;
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2658
	v2660 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2658
	v2665 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_27), int32(5))
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		goto L14
	} else {
		goto L640
	}
L597:
	;
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2461)+8))
	if v2460 < v2462 {
		v2641 = int32(0)
		goto L598
	} else {
		goto L599
	}
L598:
	;
	if v2641 == int32(0) {
		goto L596
	} else {
		goto L634
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2460
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2462
	v2469 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_28), int32(94))
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		goto L14
	} else {
		goto L601
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2465
	v2641 = v2638
	goto L598
L601:
	;
	if v2469 == int32(0) {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v2638 = int32(0)
	goto L600
L603:
	;
	goto L604
L604:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2474
	v2476 = int32(1)
	switch v2469 - v2476 {
	case 0:
		goto L606
	case 1:
		goto L605
	default:
		v2638 = v2476
		goto L600
	}
L605:
	;
	v2633 = F_slice_del(m, l0)
	mBase = m.M
	v2634 = m.ExcPending
	if v2634 != 0 {
		goto L14
	} else {
		goto L632
	}
L606:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2480 = int32(0)
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L609
L607:
	;
	if v2610 != 0 {
		goto L625
	} else {
		goto L626
	}
L608:
	;
	v2610 = v2603
	goto L607
L609:
	;
	if v2493 <= v2494 {
		v2603 = int32(-1)
		goto L608
	} else {
		goto L611
	}
L610:
	;
	v2603 = int32(0)
	goto L608
L611:
	;
	v2511 = int32(1)
	v2512 = v2493 - v2511
	v2514 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2495+v2512))))
	v2516 = v2514 & int32(255)
	if base.B2i32(v2512 == v2494)|base.B2i32(int32(0) <= v2514) != 0 {
		v2574 = v2516
		v2578 = v2511
		goto L612
	} else {
		goto L613
	}
L612:
	;
	if int32(259) < v2574 {
		goto L620
	} else {
		goto L621
	}
L613:
	;
	v2523 = v2516 & int32(63)
	v2525 = v2493 - int32(2)
	v2527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2495+v2525))))
	v2529 = v2527 << (uint(int32(6)) % 32)
	if base.B2i32(v2525 != v2494)&base.B2i32(base.Ui32(v2527) < base.Ui32(int32(192))) == int32(0) {
		goto L614
	} else {
		goto L615
	}
L614:
	;
	v2574 = v2529&int32(1984) | v2523
	v2578 = int32(2)
	goto L612
L615:
	;
	goto L616
L616:
	;
	v2542 = v2529&int32(4032) | v2523
	v2544 = v2493 - int32(3)
	v2546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2495+v2544))))
	if base.B2i32(v2544 != v2494)&base.B2i32(base.Ui32(v2546) < base.Ui32(int32(224))) == int32(0) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v2574 = v2546<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v2542
	v2578 = int32(3)
	goto L612
L618:
	;
	goto L619
L619:
	;
	v2564 = int32(4)
	v2566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2493+v2495-v2564))))
	v2574 = v2546<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_29) | v2566&int32(7)<<(uint(int32(18))%32) | v2542
	v2578 = v2564
	goto L612
L620:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2493 - v2578
	goto L624
L621:
	;
	v2580 = v2574 - int32(97)
	if v2580 < int32(0) {
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v2586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2580)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2586)>>(uint(v2580&int32(7))%32))&int32(1) == int32(0) {
		goto L620
	} else {
		goto L623
	}
L623:
	;
	v2610 = v2578
	goto L607
L624:
	;
	goto L610
L625:
	;
	v2611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2613 = v2611 + (v2474 - v2479)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2613
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2613 <= v2615 {
		v2638 = v2480
		goto L600
	} else {
		goto L628
	}
L626:
	;
	goto L627
L627:
	;
	v2629 = F_slice_del(m, l0)
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L14
	} else {
		goto L630
	}
L628:
	;
	v2617 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2617+v2613-int32(1)))))
	if v2621 != int32(117) {
		v2638 = v2480
		goto L600
	} else {
		goto L629
	}
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2613 - int32(1)
	goto L627
L630:
	;
	if int32(0) <= v2629 {
		v2638 = int32(1)
		goto L600
	} else {
		goto L631
	}
L631:
	;
	v2641 = v2629
	goto L598
L632:
	;
	if v2633 < int32(0) {
		v2641 = v2633
		goto L598
	} else {
		goto L633
	}
L633:
	;
	v2638 = v2476
	goto L600
L634:
	;
	v2647 = int32(0)
	v2648 = base.B2i32(v2641 < v2647)
	if v2648 == v2647 {
		goto L596
	} else {
		goto L635
	}
L635:
	;
	if v2641 < v2647 {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v2652 = v2641
	goto L638
L637:
	;
	v2652 = int32(1)
	goto L638
L638:
	;
	return v2652
L639:
	;
	if v2680 < int32(0) {
		v2805 = v2680
		goto L1
	} else {
		goto L647
	}
L640:
	;
	if v2665 == int32(0) {
		v2680 = v2660
		goto L639
	} else {
		goto L641
	}
L641:
	;
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2669
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2671)+8))
	if v2669 < v2672 {
		v2680 = v2660
		goto L639
	} else {
		goto L642
	}
L642:
	;
	v2675 = F_slice_del(m, l0)
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L14
	} else {
		goto L643
	}
L643:
	;
	if int32(0) <= v2675 {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	v2679 = int32(1)
	goto L646
L645:
	;
	v2679 = v2675
	goto L646
L646:
	;
	v2680 = v2679
	goto L639
L647:
	;
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2684
	goto L649
L648:
	;
	if v2791 < int32(0) {
		v2805 = v2791
		goto L1
	} else {
		goto L686
	}
L649:
	;
	v2692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2692
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2694 <= v2692 {
		goto L652
	} else {
		goto L653
	}
L650:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2692
	v2791 = int32(1)
	goto L648
L651:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L665
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2692
	v2731 = v2692
	v2732 = v2694
	goto L651
L653:
	;
	v2696 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2696+v2692))))
	v2700 = v2698 - int32(73)
	v2701 = int32(0)
	if base.B2i32(v2700 == v2701)|base.B2i32(v2700 == int32(12)) == v2701 {
		goto L652
	} else {
		goto L654
	}
L654:
	;
	v2710 = F_find_among(m, l0, int32(_a_F_romanian_UTF_8_stem_30), int32(3))
	mBase = m.M
	v2711 = m.ExcPending
	if v2711 != 0 {
		goto L14
	} else {
		goto L655
	}
L655:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2712
	switch v2710 - int32(1) {
	case 0:
		goto L657
	case 1:
		goto L656
	case 2:
		goto L658
	default:
		goto L649
	}
L656:
	;
	v2725 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_31))
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L14
	} else {
		goto L661
	}
L657:
	;
	v2719 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_32))
	mBase = m.M
	v2720 = m.ExcPending
	if v2720 != 0 {
		goto L14
	} else {
		goto L659
	}
L658:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2731 = v2712
	v2732 = v2716
	goto L651
L659:
	;
	if int32(0) <= v2719 {
		goto L649
	} else {
		goto L660
	}
L660:
	;
	v2791 = v2719
	goto L648
L661:
	;
	if int32(0) <= v2725 {
		goto L649
	} else {
		goto L662
	}
L662:
	;
	v2791 = v2725
	goto L648
L663:
	;
	if int32(0) <= v2785 {
		goto L683
	} else {
		goto L684
	}
L665:
	;
	goto L666
L666:
	;
	goto L667
L667:
	;
	v2740 = v2731
	v2742 = int32(1)
	goto L670
L669:
	;
	v2785 = v2770
	goto L663
L670:
	;
	if v2732 <= v2740 {
		goto L672
	} else {
		goto L673
	}
L671:
	;
	goto L669
L672:
	;
	v2785 = int32(-1)
	goto L663
L673:
	;
	goto L674
L674:
	;
	v2747 = v2740 + int32(1)
	v2749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2733+v2740))))
	if base.Ui32(v2749) < base.Ui32(int32(192)) {
		v2770 = v2747
		goto L675
	} else {
		goto L676
	}
L675:
	;
	v2771 = int32(1)
	if v2771 < v2742 {
		v2740 = v2770
		v2742 = v2742 - v2771
		goto L670
	} else {
		goto L682
	}
L676:
	;
	if v2732 <= v2747 {
		v2770 = v2747
		goto L675
	} else {
		goto L677
	}
L677:
	;
	v2756 = v2747
	goto L678
L678:
	;
	v2759 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2733+v2756))))
	if int32(-65) < v2759 {
		v2770 = v2756
		goto L675
	} else {
		goto L680
	}
L679:
	;
	v2770 = v2732
	goto L675
L680:
	;
	v2763 = v2756 + int32(1)
	if v2763 != v2732 {
		v2756 = v2763
		goto L678
	} else {
		goto L681
	}
L681:
	;
	goto L679
L682:
	;
	goto L671
L683:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2785
	goto L649
L684:
	;
	goto L685
L685:
	;
	goto L650
L686:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2684
	v2805 = int32(1)
	goto L1
L687:
	;
	if int32(0) <= v2800 {
		goto L43
	} else {
		goto L688
	}
L688:
	;
	goto L44
}
func F_rstacktoodeep(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_rstacktoodeep[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_rstacktoodeep[1]))
	v7 = m.G0
	v10 = v6 - (v7 - int32(1))
	v12 = v10 >> (uint(int32(31)) % 32)
	return base.B2i32(v4 < v10^v12-v12) & base.B2i32(v6 != int32(0))
}
func F_rtrim(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13857(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
