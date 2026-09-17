package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_text_concat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v5 = F_concat_internal(m, int32(_a_F_text_concat_0), int32(0), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			return v5
		}
	}
}
func F_text_format_append_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	if l3 == int32(0) {
		F_appendStringInfoString(m, l0, l1)
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			return
		}
	} else {
		if l3 < int32(0) {
			if l3 == int32(-2147483648) {
				F_errstart_cold(m, int32(21), int32(0))
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errcode(m, int32(50331778))
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_text_format_append_string_0), int32(0))
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_text_format_append_string_1), int32(_a_F_text_format_append_string_2), int32(_a_F_text_format_append_string_3))
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v14 = F_pg_mbstrlen(m, l1)
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v22 = int32(0) - l3
					v23 = v14
					F_appendStringInfoString(m, l0, l1)
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if v22 <= v23 {
							return
						} else {
							F_appendStringInfoSpaces(m, l0, v22-v23)
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			v16 = F_pg_mbstrlen(m, l1)
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if l2&int32(1) == int32(0) {
					if l3 <= v16 {
						F_appendStringInfoString(m, l0, l1)
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							return
						}
					} else {
						F_appendStringInfoSpaces(m, l0, l3-v16)
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l0, l1)
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v22 = l3
					v23 = v16
					F_appendStringInfoString(m, l0, l1)
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if v22 <= v23 {
							return
						} else {
							F_appendStringInfoSpaces(m, l0, v22-v23)
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_text_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(1)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v25 = v23 & int32(1)
			if v25 != 0 {
				v26 = v15
			} else {
				v26 = v10 + int32(4)
			}
			if v23 == int32(1) {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v32 == int32(18) {
					v35 = int32(16)
				} else {
					v35 = int32(0)
				}
				if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v42 = int32(4)
				} else {
					v42 = v35
				}
				v53 = v42
			} else {
				v43 = int32(1)
				if v25 != 0 {
					v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = int32(1)
			v55 = v17 + v54
			if v19&v54 != 0 {
				v60 = v55
			} else {
				v60 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
				if v66 == int32(18) {
					v69 = int32(16)
				} else {
					v69 = int32(0)
				}
				if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v76 = int32(4)
				} else {
					v76 = v69
				}
				v89 = v76
			} else {
				v77 = int32(1)
				if v19&v77 != 0 {
					v89 = int32(base.Ui32(v19)>>(uint(v77)%32)) - v77
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_varstr_cmp(m, v26, v53, v60, v89, v20)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v90^int32(-1)) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v90^int32(-1)) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v90^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v90^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_text_pattern_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v77 = int32(1)
			if v17&v77 != 0 {
				v81 = v77
			} else {
				v81 = int32(4)
			}
			v83 = int32(1)
			if v47&v83 != 0 {
				v87 = v83
			} else {
				v87 = int32(4)
			}
			v89 = base.B2i32(v46 < v76)
			if v46 < v76 {
				v90 = v46
			} else {
				v90 = v76
			}
			v91 = F_memcmp(m, v6+v81, v11+v87, v90)
			mBase = m.M
			if v91 != 0 {
				v94 = v91
			} else {
				if v46 < v76 {
					v94 = int32(-1)
				} else {
					v94 = base.B2i32(v76 < v46)
				}
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v95 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v99 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v94)
						}
					} else {
						return base.B2i32(int32(0) < v94)
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v99 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						return base.B2i32(int32(0) < v94)
					}
				} else {
					return base.B2i32(int32(0) < v94)
				}
			}
		}
	}
}
func F_text_regclass(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_textToQualifiedNameList(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = F_makeRangeVarFromNameList(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v11 = int32(0)
				v15 = F_RangeVarGetRelidExtended(m, v9, v11, v11, v11, v11)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					return v15
				}
			}
		}
	}
}
func F_text_starts_with(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L60
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_newlocale_from_collation(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L55
	}
L5:
	;
	return int32(0)
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v17 = F_toast_raw_datum_size(m, v9)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v20 = F_toast_raw_datum_size(m, v8)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L10
	}
L9:
	;
	return v143
L10:
	;
	if base.Ui32(v17) < base.Ui32(v20) {
		v143 = int32(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v25 = F_text_substring(m, v9, int32(1), v20, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v27 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v29 = int32(1)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v31&v29 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v34 = v29
	goto L16
L15:
	;
	v34 = int32(4)
	goto L16
L16:
	;
	v35 = v25 + v34
	v36 = int32(1)
	v37 = v27 + v36
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v42 = v40 & v36
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v43 = v37
	goto L19
L18:
	;
	v43 = v27 + int32(4)
	goto L19
L19:
	;
	if v40 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v49 == int32(18) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v60 = int32(1)
	if v42 != 0 {
		v70 = int32(base.Ui32(v40)>>(uint(v60)%32)) - v60
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v52 = int32(16)
	goto L26
L25:
	;
	v52 = int32(0)
	goto L26
L26:
	;
	if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v59 = int32(4)
	goto L29
L28:
	;
	v59 = v52
	goto L29
L29:
	;
	v70 = v59
	goto L20
L30:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v133 != v25 {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	v132 = int32(0)
	goto L31
L33:
	;
	v106 = v101
	v107 = v102
	v108 = v103
	goto L43
L34:
	;
	if (v35|v43)&int32(3) != 0 {
		v101 = v35
		v102 = v43
		v103 = v70
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v94 = v35
	v95 = v43
	v96 = v70
	goto L36
L36:
	;
	if v96 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v78 = v35
	v79 = v43
	v80 = v70
	goto L38
L38:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v83 != v84 {
		v101 = v78
		v102 = v79
		v103 = v80
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v94 = v89
	v95 = v87
	v96 = v91
	goto L36
L40:
	;
	v86 = int32(4)
	v87 = v79 + v86
	v89 = v78 + v86
	v91 = v80 - v86
	if base.Ui32(int32(3)) < base.Ui32(v91) {
		v78 = v89
		v79 = v87
		v80 = v91
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v101 = v94
	v102 = v95
	v103 = v96
	goto L33
L43:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v111 == v112 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v132 = v111 - v112
	goto L31
L45:
	;
	v114 = int32(1)
	v119 = v108 - v114
	if v119 != 0 {
		v106 = v106 + v114
		v107 = v107 + v114
		v108 = v119
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	F_pfree(m, v25)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v138 = base.B2i32(v132 == int32(0))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v27 == v139 {
		v143 = v138
		goto L9
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	F_pfree(m, v27)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v143 = v138
	goto L9
L55:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_text_starts_with_0), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	F_errhint(m, int32(_a_F_text_starts_with_1), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_text_starts_with_2), int32(1648), int32(_a_F_text_starts_with_3))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_text_starts_with_4), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_text_starts_with_2), int32(1920), int32(_a_F_text_starts_with_5))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_text_starts_with_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_like_regex_support(m, v2, int32(4))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_text_to_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(0)
	F_InitMaterializedSRF(m, l0, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v18
		v22 = F_split_text(m, l0, v6+int32(4))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return int32(0)
		}
	}
}
