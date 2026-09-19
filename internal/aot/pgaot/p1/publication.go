package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPublicationRelations(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = F_table_open(m, int32(_a_F_GetPublicationRelations_0), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v9, int32(2), int32(3), int32(184), l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(0)
	v24 = int32(1)
	v27 = F_systable_beginscan(m, v13, int32(_a_F_GetPublicationRelations_1), v24, v22, v24, v9)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = v22
	goto L5
L5:
	;
	v35 = F_systable_getnext(m, v27)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_systable_endscan(m, v27)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)+8))
	v41 = F_GetPubPartitionOptionRelations(m, v29, l1, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L6
L11:
	;
	v29 = v41
	goto L5
L12:
	;
	F_relation_close(m, v13, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_list_sort(m, v29, int32(467))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v51 = int32(0)
	if v29 == v51 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	m.G0 = v9 + int32(48)
	return v29
L16:
	;
	goto L15
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v61 < int32(2) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v64 = int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v61 != int32(2) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v146 + int32(1)
	goto L16
L20:
	;
	v68 = int32(1)
	v69 = v61 - v68
	v74 = int32(0)
	v77 = v74
	v79 = v64
	v80 = v74
	goto L23
L21:
	;
	v122 = v51
	v124 = v64
	goto L22
L22:
	;
	v130 = int32(2)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v65+v124<<(uint(v130)%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v65+v122<<(uint(v130)%32))))
	if v133 == v137 {
		v146 = v122
		goto L19
	} else {
		goto L33
	}
L23:
	;
	v85 = int32(2)
	v87 = v65 + v79<<(uint(v85)%32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65+v77<<(uint(v85)%32))))
	if v88 != v92 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v69&v68 == int32(0) {
		v146 = v113
		goto L19
	} else {
		goto L32
	}
L25:
	;
	v95 = v77 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65+v95<<(uint(int32(2))%32)))) = v88
	v100 = v95
	goto L27
L26:
	;
	v100 = v77
	goto L27
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v65+v100<<(uint(int32(2))%32))))
	if v101 != v105 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v108 = v100 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65+v108<<(uint(int32(2))%32)))) = v101
	v113 = v108
	goto L30
L29:
	;
	v113 = v100
	goto L30
L30:
	;
	v114 = int32(2)
	v115 = v79 + v114
	v117 = v80 + v114
	if v117 != v69&int32(-2) {
		v77 = v113
		v79 = v115
		v80 = v117
		goto L23
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	v122 = v113
	v124 = v115
	goto L22
L33:
	;
	v140 = v122 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65+v140<<(uint(int32(2))%32)))) = v133
	v146 = v140
	goto L19
}
func F_getPublicationSchemaInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_SearchSysCache1(m, int32(49), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			if l1 != 0 {
				v76 = int32(0)
				m.G0 = v10 + int32(32)
				return v76
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v25
					F_errmsg_internal(m, int32(_a_F_getPublicationSchemaInfo_0), v10)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_getPublicationSchemaInfo_1), int32(2876), int32(_a_F_getPublicationSchemaInfo_2))
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
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
			v37 = v35 + v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
			v39 = F_get_publication_name(m, v38, l1)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v39
				if v39 == int32(0) {
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v76 = base.B2i32(v39 != int32(0))
						m.G0 = v10 + int32(32)
						return v76
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					v45 = F_get_namespace_name(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v45
						if v45 != 0 {
							F_ReleaseCatCache(m, v14)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v76 = base.B2i32(v39 != int32(0))
								m.G0 = v10 + int32(32)
								return v76
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							F_pfree(m, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v14)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if l1 != 0 {
										v76 = int32(0)
										m.G0 = v10 + int32(32)
										return v76
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v48
											F_errmsg_internal(m, int32(_a_F_getPublicationSchemaInfo_3), v10+int32(16))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getPublicationSchemaInfo_1), int32(2897), int32(_a_F_getPublicationSchemaInfo_2))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
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
							}
						}
					}
				}
			}
		}
	}
}
func F_get_publication_name(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn13904(m, l0, l1, int32(4), int32(_a_F_get_publication_name_0), int32(3796), int32(_a_F_get_publication_name_1), int32(51))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
