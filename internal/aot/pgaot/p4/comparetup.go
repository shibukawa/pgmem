package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_comparetup_datum(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 == int32(1) {
		if v8&int32(1) != 0 {
			v33 = v7
			v36 = int32(0)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
			if v37 == v36 {
				v64 = v36
				return v64
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
				if v41 == int32(1) {
					if v40&int32(1) != 0 {
						v64 = v36
						return v64
					} else {
						v68 = v33
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+9)))
						if v73 != 0 {
							v74 = int32(-1)
						} else {
							v74 = int32(1)
						}
						return v74
					}
				} else {
					if v40&int32(1) != 0 {
						v77 = v33
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+9)))
						if v82 != 0 {
							v83 = int32(1)
						} else {
							v83 = int32(-1)
						}
						return v83
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
						v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, v48, v49, v33)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+8)))
							if v53 != int32(1) {
								v64 = v51
							} else {
								v57 = int32(0)
								if v51 < v57 {
									v61 = int32(1)
								} else {
									v61 = v57 - v51
								}
								v64 = v61
							}
							return v64
						}
					}
				}
			}
		} else {
			v68 = v7
			v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+9)))
			if v73 != 0 {
				v74 = int32(-1)
			} else {
				v74 = int32(1)
			}
			return v74
		}
	} else {
		if v8&int32(1) != 0 {
			v77 = v7
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+9)))
			if v82 != 0 {
				v83 = int32(1)
			} else {
				v83 = int32(-1)
			}
			return v83
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int32)(m, v16, v17, v7)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = int32(1)
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)))
				if v24 != v23 {
					v31 = v19
					if v31 != 0 {
						v64 = v31
						return v64
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
						v33 = v32
						v36 = int32(0)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
						if v37 == v36 {
							v64 = v36
							return v64
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
							if v41 == int32(1) {
								if v40&int32(1) != 0 {
									v64 = v36
									return v64
								} else {
									v68 = v33
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+9)))
									if v73 != 0 {
										v74 = int32(-1)
									} else {
										v74 = int32(1)
									}
									return v74
								}
							} else {
								if v40&int32(1) != 0 {
									v77 = v33
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+9)))
									if v82 != 0 {
										v83 = int32(1)
									} else {
										v83 = int32(-1)
									}
									return v83
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
									v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, v48, v49, v33)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+8)))
										if v53 != int32(1) {
											v64 = v51
										} else {
											v57 = int32(0)
											if v51 < v57 {
												v61 = int32(1)
											} else {
												v61 = v57 - v51
											}
											v64 = v61
										}
										return v64
									}
								}
							}
						}
					}
				} else {
					if v19 < int32(0) {
						v64 = v23
						return v64
					} else {
						v31 = int32(0) - v19
						if v31 != 0 {
							v64 = v31
							return v64
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
							v33 = v32
							v36 = int32(0)
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
							if v37 == v36 {
								v64 = v36
								return v64
							} else {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
								if v41 == int32(1) {
									if v40&int32(1) != 0 {
										v64 = v36
										return v64
									} else {
										v68 = v33
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+9)))
										if v73 != 0 {
											v74 = int32(-1)
										} else {
											v74 = int32(1)
										}
										return v74
									}
								} else {
									if v40&int32(1) != 0 {
										v77 = v33
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+9)))
										if v82 != 0 {
											v83 = int32(1)
										} else {
											v83 = int32(-1)
										}
										return v83
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
										v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, v48, v49, v33)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+8)))
											if v53 != int32(1) {
												v64 = v51
											} else {
												v57 = int32(0)
												if v51 < v57 {
													v61 = int32(1)
												} else {
													v61 = v57 - v51
												}
												v64 = v61
											}
											return v64
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
func F_comparetup_datum_tiebreak(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	if v7 == v4 {
		v44 = v4
		return v44
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v11 == int32(1) {
			if v10&int32(1) != 0 {
				v44 = v4
				return v44
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)))
				if v18 != 0 {
					v19 = int32(-1)
				} else {
					v19 = int32(1)
				}
				return v19
			}
		} else {
			if v10&int32(1) != 0 {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)))
				if v25 != 0 {
					v26 = int32(1)
				} else {
					v26 = int32(-1)
				}
				return v26
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
				v31 = m.T0[v30].(func(*base.Module, int32, int32, int32) int32)(m, v28, v29, v6)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)))
					if v35 != int32(1) {
						v44 = v31
					} else {
						v39 = int32(0)
						if v31 < v39 {
							v43 = int32(1)
						} else {
							v43 = v39 - v31
						}
						v44 = v43
					}
					return v44
				}
			}
		}
	}
}
func F_comparetup_index_btree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 == int32(1) {
		if v7&int32(1) != 0 {
			v44 = F_comparetup_index_btree_tiebreak(m, l0, l1, l2)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v47 = v44
				return v47
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)))
			if v16 != 0 {
				v17 = int32(-1)
			} else {
				v17 = int32(1)
			}
			return v17
		}
	} else {
		if v7&int32(1) != 0 {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)))
			if v23 != 0 {
				v24 = int32(1)
			} else {
				v24 = int32(-1)
			}
			return v24
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, v26, v27, v8)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = int32(1)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)))
				if v34 != v33 {
					v41 = v29
					if v41 != 0 {
						v47 = v41
						return v47
					} else {
						v44 = F_comparetup_index_btree_tiebreak(m, l0, l1, l2)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v47 = v44
							return v47
						}
					}
				} else {
					if v29 < int32(0) {
						v47 = v33
						return v47
					} else {
						v41 = int32(0) - v29
						if v41 != 0 {
							v47 = v41
							return v47
						} else {
							v44 = F_comparetup_index_btree_tiebreak(m, l0, l1, l2)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v47 = v44
								return v47
							}
						}
					}
				}
			}
		}
	}
}
func F_comparetup_index_btree_tiebreak(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	if v25 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(208)
	return v222
L2:
	;
	v70 = int32(2)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v70 <= v18 {
		goto L27
	} else {
		goto L28
	}
L3:
	;
	v31 = F_index_getattr_2(m, v20, int32(1), v23, v16+int32(207))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v38 = F_index_getattr_2(m, v19, int32(1), v23, v16+int32(206))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+206)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+207)))
	if v41 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v40&int32(1) != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v40&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	if v48 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = int32(-1)
	goto L13
L12:
	;
	v49 = int32(1)
	goto L13
L13:
	;
	v222 = v49
	goto L1
L14:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v57 = m.T0[v56].(func(*base.Module, int32, int32, int32) int32)(m, v31, v38, v24)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L20
	}
L17:
	;
	v55 = int32(1)
	goto L19
L18:
	;
	v55 = int32(-1)
	goto L19
L19:
	;
	v222 = v55
	goto L1
L20:
	;
	v59 = int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
	if v60 != v59 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = v57
	goto L23
L22:
	;
	if v57 < int32(0) {
		v222 = v59
		goto L1
	} else {
		goto L24
	}
L23:
	;
	if v67 != 0 {
		v222 = v67
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v67 = int32(0) - v57
	goto L23
L25:
	;
	goto L2
L26:
	;
	v222 = int32(1)
	goto L1
L27:
	;
	v75 = v24
	v76 = v70
	v83 = v71
	goto L30
L28:
	;
	v139 = v71
	goto L29
L29:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
	if v143 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L30:
	;
	v88 = v75 + int32(36)
	v91 = F_index_getattr_2(m, v20, v76, v23, v16+int32(207))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	v139 = v126
	goto L29
L32:
	;
	v95 = F_index_getattr_2(m, v19, v76, v23, v16+int32(206))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+206)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+207)))
	if v98 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v128 = v76 + int32(1)
	if v128 <= v18 {
		v75 = v88
		v76 = v128
		v83 = v126
		goto L30
	} else {
		goto L54
	}
L35:
	;
	v101 = int32(1)
	if v97&v101 != 0 {
		v126 = v101
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v97&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+45)))
	if v106 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v107 = int32(-1)
	goto L41
L40:
	;
	v107 = int32(1)
	goto L41
L41:
	;
	v222 = v107
	goto L1
L42:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+45)))
	if v112 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v75)+52))
	v115 = m.T0[v114].(func(*base.Module, int32, int32, int32) int32)(m, v91, v95, v88)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L48
	}
L45:
	;
	v113 = int32(1)
	goto L47
L46:
	;
	v113 = int32(-1)
	goto L47
L47:
	;
	v222 = v113
	goto L1
L48:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+44)))
	if v117 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v115 < int32(0) {
		goto L26
	} else {
		goto L52
	}
L50:
	;
	v124 = v115
	goto L51
L51:
	;
	if v124 != 0 {
		v222 = v124
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v124 = int32(0) - v115
	goto L51
L53:
	;
	v126 = v83
	goto L34
L54:
	;
	goto L31
L55:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+2)))
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20))))
	v200 = int32(16)
	v202 = v198 | v199<<(uint(v200)%32)
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+2)))
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19))))
	v207 = v203 | v204<<(uint(v200)%32)
	if v202 != v207 {
		goto L71
	} else {
		goto L72
	}
L56:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+9)))
	if (v146^int32(-1))&v139&int32(1) != 0 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v153 = v16 - int32(-64)
	v155 = v16 + int32(32)
	F_index_deform_tuple(m, v20, v23, v153, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v159 = F_BuildIndexValueDescription(m, v158, v153, v155)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(83906754))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v169 + int32(4)
	F_errmsg(m, int32(_a_F_comparetup_index_btree_tiebreak_0), v16+int32(16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	if v159 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+48))
	F_errtableconstraint(m, v186, v188+int32(4))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L69
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v159
	F_errdetail(m, int32(_a_F_comparetup_index_btree_tiebreak_1), v16)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	F_errdetail(m, int32(_a_F_comparetup_index_btree_tiebreak_2), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L68
	}
L67:
	;
	goto L63
L68:
	;
	goto L63
L69:
	;
	F_errfinish(m, int32(_a_F_comparetup_index_btree_tiebreak_3), int32(1693), int32(_a_F_comparetup_index_btree_tiebreak_4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	if base.Ui32(v202) < base.Ui32(v207) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v222 = base.B2i32(base.Ui32(v214) < base.Ui32(v213)) - base.B2i32(base.Ui32(v213) < base.Ui32(v214))
	goto L1
L74:
	;
	v212 = int32(-1)
	goto L76
L75:
	;
	v212 = int32(1)
	goto L76
L76:
	;
	v222 = v212
	goto L1
}
