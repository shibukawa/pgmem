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
	v5 = F_concat_internal(m, int32(757756), int32(0), l0)
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
						F_errmsg(m, int32(402566), int32(0))
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(500555), int32(6372), int32(330295))
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
				v29 = int32(4)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v31&int32(254) == int32(2) {
					v40 = v29
				} else {
					v40 = base.B2i32(v31 == int32(18)) << (uint(v29) % 32)
				}
				if v31 == int32(1) {
					v43 = v29
				} else {
					v43 = v40
				}
				v54 = v43
			} else {
				v44 = int32(1)
				if v25 != 0 {
					v54 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v55 = int32(1)
			v56 = v17 + v55
			if v19&v55 != 0 {
				v61 = v56
			} else {
				v61 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v64 = int32(4)
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				if v66&int32(254) == int32(2) {
					v75 = v64
				} else {
					v75 = base.B2i32(v66 == int32(18)) << (uint(v64) % 32)
				}
				if v66 == int32(1) {
					v78 = v64
				} else {
					v78 = v75
				}
				v91 = v78
			} else {
				v79 = int32(1)
				if v19&v79 != 0 {
					v91 = int32(base.Ui32(v19)>>(uint(v79)%32)) - v79
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v92 = F_varstr_cmp(m, v26, v54, v61, v91, v20)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v94 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v98 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v92^int32(-1)) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v92^int32(-1)) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v98 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v92^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v92^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_text_pattern_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if v20 == int32(1) {
				v23 = int32(4)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
				if v25&int32(254) == int32(2) {
					v34 = v23
				} else {
					v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
				}
				if v25 == int32(1) {
					v37 = v23
				} else {
					v37 = v34
				}
				v50 = v37
			} else {
				v38 = int32(1)
				if v20&v38 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v38)%32)) - v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v51 == int32(1) {
				v54 = int32(4)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v56&int32(254) == int32(2) {
					v65 = v54
				} else {
					v65 = base.B2i32(v56 == int32(18)) << (uint(v54) % 32)
				}
				if v56 == int32(1) {
					v68 = v54
				} else {
					v68 = v65
				}
				v81 = v68
			} else {
				v69 = int32(1)
				if v51&v69 != 0 {
					v81 = int32(base.Ui32(v51)>>(uint(v69)%32)) - v69
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v82 = int32(1)
			if v20&v82 != 0 {
				v86 = v82
			} else {
				v86 = int32(4)
			}
			v88 = int32(1)
			if v51&v88 != 0 {
				v92 = v88
			} else {
				v92 = int32(4)
			}
			v94 = base.B2i32(v50 < v81)
			if v50 < v81 {
				v95 = v50
			} else {
				v95 = v81
			}
			v96 = F_memcmp(m, v7+v86, v14+v92, v95)
			mBase = m.M
			if v96 != 0 {
				v99 = v96
			} else {
				if v50 < v81 {
					v99 = int32(-1)
				} else {
					v99 = base.B2i32(v81 < v50)
				}
			}
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v100 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v104 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v99)
						}
					} else {
						return base.B2i32(int32(0) < v99)
					}
				}
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v104 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						return base.B2i32(int32(0) < v99)
					}
				} else {
					return base.B2i32(int32(0) < v99)
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
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
	v173 = m.ExcPending
	if v173 != 0 {
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
	v153 = m.ExcPending
	if v153 != 0 {
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
	return v144
L10:
	;
	if base.Ui32(v17) < base.Ui32(v20) {
		v144 = int32(0)
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
	if base.Ui32(int32(4)) <= base.Ui32(v71) {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v46 = int32(4)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v48&int32(254) == int32(2) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v61 = int32(1)
	if v42 != 0 {
		v71 = int32(base.Ui32(v40)>>(uint(v61)%32)) - v61
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v57 = v46
	goto L26
L25:
	;
	v57 = base.B2i32(v48 == int32(18)) << (uint(v46) % 32)
	goto L26
L26:
	;
	if v48 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v60 = v46
	goto L29
L28:
	;
	v60 = v57
	goto L29
L29:
	;
	v71 = v60
	goto L20
L30:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v134 != v25 {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	v133 = int32(0)
	goto L31
L33:
	;
	v107 = v102
	v108 = v103
	v109 = v104
	goto L43
L34:
	;
	if (v35|v43)&int32(3) != 0 {
		v102 = v35
		v103 = v43
		v104 = v71
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v95 = v35
	v96 = v43
	v97 = v71
	goto L36
L36:
	;
	if v97 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v79 = v35
	v80 = v43
	v81 = v71
	goto L38
L38:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v84 != v85 {
		v102 = v79
		v103 = v80
		v104 = v81
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v95 = v90
	v96 = v88
	v97 = v92
	goto L36
L40:
	;
	v87 = int32(4)
	v88 = v80 + v87
	v90 = v79 + v87
	v92 = v81 - v87
	if base.Ui32(int32(3)) < base.Ui32(v92) {
		v79 = v90
		v80 = v88
		v81 = v92
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v102 = v95
	v103 = v96
	v104 = v97
	goto L33
L43:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v112 == v113 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v133 = v112 - v113
	goto L31
L45:
	;
	v115 = int32(1)
	v120 = v109 - v115
	if v120 != 0 {
		v107 = v107 + v115
		v108 = v108 + v115
		v109 = v120
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
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v139 = base.B2i32(v133 == int32(0))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v27 == v140 {
		v144 = v139
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
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v144 = v139
	goto L9
L55:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(245604), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	F_errhint(m, int32(575250), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(500555), int32(1648), int32(106279))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
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
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(169141), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(500555), int32(1920), int32(320743))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
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
